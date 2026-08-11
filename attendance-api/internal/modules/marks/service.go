package marks

import (
	"context"
	"fmt"

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
	isAssigned, err := s.repo.IsFacultyAssignedRightNow(ctx, facultyID, req.ClassID, req.SubjectID)
	if err != nil {
		return err
	}
	if !isAssigned {
		return fmt.Errorf("faculty is not assigned to this class and subject right now according to the timetable")
	}

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
