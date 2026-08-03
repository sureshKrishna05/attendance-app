package marks

import (
	"context"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

type Repository interface {
	CreateExam(ctx context.Context, exam *Exam) error
	UploadMarks(ctx context.Context, marks []Mark) error
	GetStudentMarks(ctx context.Context, studentID string) ([]StudentMarkResult, error)
}

type repository struct {
	db *pgxpool.Pool
}

func NewRepository(db *pgxpool.Pool) Repository {
	return &repository{db: db}
}

func (r *repository) CreateExam(ctx context.Context, e *Exam) error {
	query := `
		INSERT INTO internal_exams (id, class_id, subject_id, name, max_marks, date)
		VALUES ($1, $2, $3, $4, $5, $6)
	`
	_, err := r.db.Exec(ctx, query, e.ID, e.ClassID, e.SubjectID, e.Name, e.MaxMarks, e.Date)
	return err
}

func (r *repository) UploadMarks(ctx context.Context, marks []Mark) error {
	batch := &pgx.Batch{}
	for _, m := range marks {
		query := `
			INSERT INTO internal_marks (exam_id, student_id, marks_obtained) 
			VALUES ($1, $2, $3)
			ON CONFLICT (exam_id, student_id) DO UPDATE SET marks_obtained = EXCLUDED.marks_obtained
		`
		batch.Queue(query, m.ExamID, m.StudentID, m.MarksObtained)
	}

	br := r.db.SendBatch(ctx, batch)
	defer br.Close()

	for i := 0; i < len(marks); i++ {
		if _, err := br.Exec(); err != nil {
			return err
		}
	}
	return nil
}

func (r *repository) GetStudentMarks(ctx context.Context, studentID string) ([]StudentMarkResult, error) {
	query := `
		SELECT e.subject_id, e.name, m.marks_obtained, e.max_marks, CAST(e.date AS TEXT)
		FROM internal_marks m
		JOIN internal_exams e ON m.exam_id = e.id
		WHERE m.student_id = $1
		ORDER BY e.date DESC
	`
	rows, err := r.db.Query(ctx, query, studentID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var results []StudentMarkResult
	for rows.Next() {
		var res StudentMarkResult
		if err := rows.Scan(&res.SubjectID, &res.ExamName, &res.MarksObtained, &res.MaxMarks, &res.Date); err != nil {
			return nil, err
		}
		results = append(results, res)
	}
	return results, nil
}
