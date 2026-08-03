package users

import (
	"context"
	"errors"

	"golang.org/x/crypto/bcrypt"
)

type Service interface {
	VerifyCredentials(ctx context.Context, id, password, role string) (*User, error)
	RegisterUser(ctx context.Context, id, password, role, name string) error
}

type service struct {
	repo Repository
}

func NewService(repo Repository) Service {
	return &service{repo: repo}
}

func (s *service) VerifyCredentials(ctx context.Context, id, password, role string) (*User, error) {
	user, err := s.repo.FindByID(ctx, id)
	if err != nil {
		return nil, err
	}
	if user == nil {
		return nil, errors.New("user not found")
	}

	if user.Role != role {
		return nil, errors.New("invalid role")
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(password))
	if err != nil {
		return nil, errors.New("invalid credentials")
	}

	return user, nil
}

func (s *service) RegisterUser(ctx context.Context, id, password, role, name string) error {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	user := &User{
		ID:           id,
		PasswordHash: string(hashedPassword),
		Role:         role,
		Name:         name,
	}

	return s.repo.Create(ctx, user)
}
