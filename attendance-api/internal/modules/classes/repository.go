package classes

import (
	"context"

	"github.com/jackc/pgx/v5/pgxpool"
)

type Repository interface {
	GetFacultyAssignments(ctx context.Context, facultyID string) ([]FacultyAssignment, error)
}

type repository struct {
	db *pgxpool.Pool
}

func NewRepository(db *pgxpool.Pool) Repository {
	return &repository{db: db}
}

func (r *repository) GetFacultyAssignments(ctx context.Context, facultyID string) ([]FacultyAssignment, error) {
	query := `
		SELECT fa.class_id, fa.subject_id, c.name, c.semester, c.section
		FROM faculty_assignments fa
		JOIN classes c ON fa.class_id = c.id
		WHERE fa.faculty_id = $1
	`
	rows, err := r.db.Query(ctx, query, facultyID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var assignments []FacultyAssignment
	for rows.Next() {
		var a FacultyAssignment
		if err := rows.Scan(&a.ClassID, &a.SubjectID, &a.ClassName, &a.Semester, &a.Section); err != nil {
			return nil, err
		}
		assignments = append(assignments, a)
	}
	return assignments, nil
}
