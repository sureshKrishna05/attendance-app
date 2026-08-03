package marks

import (
	"context"

	"github.com/google/uuid"
)

type Service interface {
	UploadMarks(ctx context.Context, facultyID string, req UploadMarksRequest) error
	GetStudentMarks(ctx context.Context, studentID string) ([]StudentMarkResult, error)
}

type service struct {
	repo Repository
}

func NewService(repo Repository) Service {
	return &service{repo: repo}
}

func (s *service) UploadMarks(ctx context.Context, facultyID string, req UploadMarksRequest) error {
	examID := uuid.New().String()

	exam := &Exam{
		ID:        examID,
		ClassID:   req.ClassID,
		SubjectID: req.SubjectID,
		Name:      req.ExamName,
		MaxMarks:  req.MaxMarks,
		Date:      req.Date,
	}

	if err := s.repo.CreateExam(ctx, exam); err != nil {
		return err
	}

	for i := range req.Marks {
		req.Marks[i].ExamID = examID
	}

	return s.repo.UploadMarks(ctx, req.Marks)
}

func (s *service) GetStudentMarks(ctx context.Context, studentID string) ([]StudentMarkResult, error) {
	return s.repo.GetStudentMarks(ctx, studentID)
}
