package user

import (
	"context"
	"errors"
	"go-auth/internal/auth"
	"strings"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"golang.org/x/crypto/bcrypt"
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

type LoginInput struct {
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

	_, err := s.repo.FindByEmail(ctx, email)

	if err == nil {
		return AuthResult{}, errors.New("email already in use")
	}

	if err != nil && !errors.Is(err, mongo.ErrNoDocuments) {
		return AuthResult{}, err
	}

	hashedPass, err := bcrypt.GenerateFromPassword([]byte(pass), bcrypt.DefaultCost)

	if err != nil {
		return AuthResult{}, errors.New("failed to hash password")
	}

	now := time.Now().UTC()

	u := User{
		Email:        email,
		PasswordHash: string(hashedPass),
		Role:         "user",
		CreatedAt:    now,
		UpdatedAt:    now,
	}

	created, err := s.repo.Create(ctx, &u)

	if err != nil {
		return AuthResult{}, err
	}

	token, err := auth.CreateToken(s.jwtSecret, created.ID.Hex(), created.Role)

	if err != nil {
		return AuthResult{}, err
	}

	return AuthResult{
		Token: token,
		User:  ToPublic(&created),
	}, nil
}

func (s *Service) Login(ctx context.Context, input LoginInput) (AuthResult, error) {
	email := strings.ToLower(strings.TrimSpace(input.Email))

	pass := strings.TrimSpace(input.Password)

	if email == "" || pass == "" {
		return AuthResult{}, errors.New("email and password are required")
	}

	if len(pass) < 6 {
		return AuthResult{}, errors.New("Password must be at least 6")
	}

	user, err := s.repo.FindByEmail(ctx,email)

	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments){
			return AuthResult{}, errors.New("Invalid Credentials")
		}
		return AuthResult{}, err
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(pass)); err != nil {
		return AuthResult{}, errors.New("Invalid Credentials")
	}

	token, err := auth.CreateToken(s.jwtSecret, user.ID.Hex(), user.Role)

	if err != nil {
		return AuthResult{}, err
	}

	return AuthResult{
		Token: token,
		User:  ToPublic(&user),
	}, nil
}