package classes

import (
	"context"
)

type Service interface {
	GetFacultyAssignments(ctx context.Context, facultyID string) ([]FacultyAssignment, error)
}

type service struct {
	repo Repository
}

func NewService(repo Repository) Service {
	return &service{repo: repo}
}

func (s *service) GetFacultyAssignments(ctx context.Context, facultyID string) ([]FacultyAssignment, error) {
	return s.repo.GetFacultyAssignments(ctx, facultyID)
}
