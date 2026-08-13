package admin

import (
	"context"

	"github.com/jackc/pgx/v5/pgxpool"
)

type Repository interface {
	BulkInsertUsers(ctx context.Context, users []BulkUser) error
	CreateClass(ctx context.Context, class *Class) error
	InsertTimetableSlot(ctx context.Context, slot *TimetableSlot) error
	CreateFacultyAssignment(ctx context.Context, assignment *FacultyAssignment) error
}

type repository struct {
	db *pgxpool.Pool
}

func NewRepository(db *pgxpool.Pool) Repository {
	return &repository{db: db}
}

func (r *repository) BulkInsertUsers(ctx context.Context, users []BulkUser) error {
	tx, err := r.db.Begin(ctx)
	if err != nil {
		return err
	}
	defer tx.Rollback(ctx)

	for _, u := range users {
		_, err := tx.Exec(ctx, "INSERT INTO users (id, password_hash, role, name, created_at, updated_at) VALUES ($1, $2, $3, $4, NOW(), NOW())", u.ID, u.Password, u.Role, u.Name)
		if err != nil {
			return err
		}
	}
	return tx.Commit(ctx)
}

func (r *repository) CreateClass(ctx context.Context, c *Class) error {
	_, err := r.db.Exec(ctx, "INSERT INTO classes (id, name, semester, section) VALUES ($1, $2, $3, $4)", c.ID, c.Name, c.Semester, c.Section)
	return err
}

func (r *repository) InsertTimetableSlot(ctx context.Context, s *TimetableSlot) error {
	_, err := r.db.Exec(ctx, "INSERT INTO timetable_slots (id, class_id, day_of_week, start_time, end_time, subject_id, faculty_id, room) VALUES ($1, $2, $3, $4, $5, $6, $7, $8)", s.ID, s.ClassID, s.DayOfWeek, s.StartTime, s.EndTime, s.SubjectID, s.FacultyID, s.Room)
	return err
}

func (r *repository) CreateFacultyAssignment(ctx context.Context, a *FacultyAssignment) error {
	_, err := r.db.Exec(ctx, "INSERT INTO faculty_assignments (class_id, subject_id, faculty_id) VALUES ($1, $2, $3)", a.ClassID, a.SubjectID, a.FacultyID)
	return err
}
