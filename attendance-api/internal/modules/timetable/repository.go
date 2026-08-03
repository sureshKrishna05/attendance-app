package timetable

import (
	"context"

	"github.com/jackc/pgx/v5/pgxpool"
)

type Repository interface {
	GetStudentTimetable(ctx context.Context, studentID string) ([]Slot, error)
	GetFacultyTimetable(ctx context.Context, facultyID string) ([]Slot, error)
}

type repository struct {
	db *pgxpool.Pool
}

func NewRepository(db *pgxpool.Pool) Repository {
	return &repository{db: db}
}

func (r *repository) GetStudentTimetable(ctx context.Context, studentID string) ([]Slot, error) {
	query := `
		SELECT ts.id, ts.class_id, ts.day_of_week, CAST(ts.start_time AS TEXT), CAST(ts.end_time AS TEXT), 
		       ts.subject_id, ts.faculty_id, u.name, ts.room
		FROM timetable_slots ts
		JOIN class_students cs ON ts.class_id = cs.class_id
		JOIN users u ON ts.faculty_id = u.id
		WHERE cs.student_id = $1
		ORDER BY ts.day_of_week, ts.start_time
	`
	rows, err := r.db.Query(ctx, query, studentID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var slots []Slot
	for rows.Next() {
		var s Slot
		if err := rows.Scan(&s.ID, &s.ClassID, &s.DayOfWeek, &s.StartTime, &s.EndTime, &s.SubjectID, &s.FacultyID, &s.FacultyName, &s.Room); err != nil {
			return nil, err
		}
		slots = append(slots, s)
	}
	return slots, nil
}

func (r *repository) GetFacultyTimetable(ctx context.Context, facultyID string) ([]Slot, error) {
	query := `
		SELECT ts.id, ts.class_id, ts.day_of_week, CAST(ts.start_time AS TEXT), CAST(ts.end_time AS TEXT), 
		       ts.subject_id, ts.faculty_id, u.name, ts.room
		FROM timetable_slots ts
		JOIN users u ON ts.faculty_id = u.id
		WHERE ts.faculty_id = $1
		ORDER BY ts.day_of_week, ts.start_time
	`
	rows, err := r.db.Query(ctx, query, facultyID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var slots []Slot
	for rows.Next() {
		var s Slot
		if err := rows.Scan(&s.ID, &s.ClassID, &s.DayOfWeek, &s.StartTime, &s.EndTime, &s.SubjectID, &s.FacultyID, &s.FacultyName, &s.Room); err != nil {
			return nil, err
		}
		slots = append(slots, s)
	}
	return slots, nil
}
