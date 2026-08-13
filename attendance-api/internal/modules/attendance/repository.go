package attendance

import (
	"context"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

type Repository interface {
	CreateSession(ctx context.Context, session *Session) error
	CreateRecords(ctx context.Context, records []Record) error
	GetStudentAttendance(ctx context.Context, studentID string) ([]StudentAttendanceResult, error)
	IsFacultyAssignedRightNow(ctx context.Context, facultyID, classID, subjectID string) (bool, error)
}

type repository struct {
	db *pgxpool.Pool
}

func NewRepository(db *pgxpool.Pool) Repository {
	return &repository{db: db}
}

func (r *repository) CreateSession(ctx context.Context, s *Session) error {
	query := `
		INSERT INTO attendance_sessions (id, subject_id, class_id, faculty_id, date, created_at)
		VALUES ($1, $2, $3, $4, $5, NOW())
	`
	_, err := r.db.Exec(ctx, query, s.ID, s.SubjectID, s.ClassID, s.FacultyID, s.Date)
	return err
}

func (r *repository) CreateRecords(ctx context.Context, records []Record) error {
	// Bulk insert using pgx batch
	batch := &pgx.Batch{}
	for _, rec := range records {
		query := `INSERT INTO attendance_records (session_id, student_id, is_present) VALUES ($1, $2, $3)`
		batch.Queue(query, rec.SessionID, rec.StudentID, rec.IsPresent)
	}
	
	br := r.db.SendBatch(ctx, batch)
	defer br.Close()
	
	for i := 0; i < len(records); i++ {
		if _, err := br.Exec(); err != nil {
			return err
		}
	}
	return nil
}

func (r *repository) GetStudentAttendance(ctx context.Context, studentID string) ([]StudentAttendanceResult, error) {
	query := `
		SELECT s.subject_id, CAST(s.date AS TEXT), r.is_present 
		FROM attendance_records r
		JOIN attendance_sessions s ON r.session_id = s.id
		WHERE r.student_id = $1
		ORDER BY s.date DESC
	`
	rows, err := r.db.Query(ctx, query, studentID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var records []StudentAttendanceResult
	for rows.Next() {
		var rec StudentAttendanceResult
		if err := rows.Scan(&rec.SubjectID, &rec.Date, &rec.IsPresent); err != nil {
			return nil, err
		}
		records = append(records, rec)
	}
	return records, nil
}

func (r *repository) IsFacultyAssignedRightNow(ctx context.Context, facultyID, classID, subjectID string) (bool, error) {
	query := `
		SELECT EXISTS (
			SELECT 1 FROM timetable_slots 
			WHERE faculty_id = $1 AND class_id = $2 AND subject_id = $3
			AND day_of_week = trim(to_char(CURRENT_DATE, 'Day'))
			AND CURRENT_TIME BETWEEN start_time AND end_time
		)
	`
	var exists bool
	err := r.db.QueryRow(ctx, query, facultyID, classID, subjectID).Scan(&exists)
	return exists, err
}
