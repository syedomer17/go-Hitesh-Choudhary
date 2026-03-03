package user

import (
	"context"
	"errors"
	"strings"

	"go.mongodb.org/mongo-driver/mongo"
)

type Service struct {
	repo *Repo

	jwtSecret string
}

func NewService(repo *Repo, jwtSecret string) *Service {
	return &Service{
		repo:      repo,
		jwtSecret: jwtSecret,
	}
}

type RegisterInput struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type AuthResult struct {
	Token string     `json:"token"`
	User  PublicUser `json:"user"`
}

func (s *Service) Register(ctx context.Context, input RegisterInput) (AuthResult, error) {
	email := strings.ToLower(strings.TrimSpace(input.Email))

	pass := strings.TrimSpace(input.Password)

	if email == "" || pass == "" {
		return AuthResult{}, errors.New("email and password are required")
	}

	if len(pass) < 6 {
		return AuthResult{}, errors.New("Password must be at least 6")
	}

	_, err := s.repo.FindByEmail(ctx,email)

	if err == nil {
		return AuthResult{}, errors.New("email already in use")
	}

	if err != nil && !errors.Is(err, mongo.ErrNoDocuments){
		return AuthResult{}, err
	}

	return AuthResult{}, nil
}
