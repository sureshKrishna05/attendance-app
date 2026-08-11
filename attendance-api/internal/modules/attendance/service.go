package attendance

import (
	"context"
	"fmt"
	"time"

	"github.com/google/uuid"
)

type Service interface {
	TakeAttendance(ctx context.Context, facultyID string, req TakeAttendanceRequest) error
	GetStudentAttendance(ctx context.Context, studentID string) ([]Record, error)
}

type service struct {
	repo Repository
}

func NewService(repo Repository) Service {
	return &service{repo: repo}
}

func (s *service) TakeAttendance(ctx context.Context, facultyID string, req TakeAttendanceRequest) error {
	isAssigned, err := s.repo.IsFacultyAssignedRightNow(ctx, facultyID, req.ClassID, req.SubjectID)
	if err != nil {
		return err
	}
	if !isAssigned {
		return fmt.Errorf("faculty is not assigned to this class and subject right now according to the timetable")
	}

	date, err := time.Parse("2006-01-02", req.Date)
	if err != nil {
		return err
	}

	sessionID := uuid.New().String()

	session := &Session{
		ID:        sessionID,
		SubjectID: req.SubjectID,
		ClassID:   req.ClassID,
		FacultyID: facultyID,
		Date:      date,
	}

	if err := s.repo.CreateSession(ctx, session); err != nil {
		return err
	}

	// Attach session ID to records
	for i := range req.Records {
		req.Records[i].SessionID = sessionID
	}

	return s.repo.CreateRecords(ctx, req.Records)
}

func (s *service) GetStudentAttendance(ctx context.Context, studentID string) ([]Record, error) {
	return s.repo.GetStudentAttendance(ctx, studentID)
}
