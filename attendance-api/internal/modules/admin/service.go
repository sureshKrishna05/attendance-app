package admin

import (
	"context"

	"golang.org/x/crypto/bcrypt"
)

type Service interface {
	BulkInsertUsers(ctx context.Context, users []BulkUser) error
	CreateClass(ctx context.Context, class *Class) error
	InsertTimetableSlot(ctx context.Context, slot *TimetableSlot) error
	CreateFacultyAssignment(ctx context.Context, assignment *FacultyAssignment) error
}

type service struct {
	repo Repository
}

func NewService(repo Repository) Service {
	return &service{repo: repo}
}

func (s *service) BulkInsertUsers(ctx context.Context, users []BulkUser) error {
	for i := range users {
		hashed, err := bcrypt.GenerateFromPassword([]byte(users[i].Password), bcrypt.DefaultCost)
		if err != nil {
			return err
		}
		users[i].Password = string(hashed) // We hijacked the Password field to store the hash for the repo
	}
	return s.repo.BulkInsertUsers(ctx, users)
}

func (s *service) CreateClass(ctx context.Context, c *Class) error {
	return s.repo.CreateClass(ctx, c)
}

func (s *service) InsertTimetableSlot(ctx context.Context, slot *TimetableSlot) error {
	return s.repo.InsertTimetableSlot(ctx, slot)
}

func (s *service) CreateFacultyAssignment(ctx context.Context, a *FacultyAssignment) error {
	return s.repo.CreateFacultyAssignment(ctx, a)
}
