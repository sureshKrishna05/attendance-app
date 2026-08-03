package attendance

import (
	"context"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

type Repository interface {
	CreateSession(ctx context.Context, session *Session) error
	CreateRecords(ctx context.Context, records []Record) error
	GetStudentAttendance(ctx context.Context, studentID string) ([]Record, error)
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

func (r *repository) GetStudentAttendance(ctx context.Context, studentID string) ([]Record, error) {
	query := `SELECT session_id, student_id, is_present FROM attendance_records WHERE student_id = $1`
	rows, err := r.db.Query(ctx, query, studentID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var records []Record
	for rows.Next() {
		var rec Record
		if err := rows.Scan(&rec.SessionID, &rec.StudentID, &rec.IsPresent); err != nil {
			return nil, err
		}
		records = append(records, rec)
	}
	return records, nil
}
