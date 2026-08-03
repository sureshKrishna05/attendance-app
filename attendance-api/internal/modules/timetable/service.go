package timetable

import (
	"context"
)

type Service interface {
	GetStudentTimetable(ctx context.Context, studentID string) ([]Slot, error)
	GetFacultyTimetable(ctx context.Context, facultyID string) ([]Slot, error)
}

type service struct {
	repo Repository
}

func NewService(repo Repository) Service {
	return &service{repo: repo}
}

func (s *service) GetStudentTimetable(ctx context.Context, studentID string) ([]Slot, error) {
	return s.repo.GetStudentTimetable(ctx, studentID)
}

func (s *service) GetFacultyTimetable(ctx context.Context, facultyID string) ([]Slot, error) {
	return s.repo.GetFacultyTimetable(ctx, facultyID)
}
