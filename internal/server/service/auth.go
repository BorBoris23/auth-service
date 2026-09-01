package service

import (
	"context"
	"errors"

	"auth-service/internal/repository"
	"auth-service/internal/server/token"

	"golang.org/x/crypto/bcrypt"
)

type AuthService struct {
	userRepository *repository.UserRepository
	jwtService     *token.JWTService
}

func NewAuthService(
	userRepository *repository.UserRepository,
	jwtService *token.JWTService,
) *AuthService {
	return &AuthService{
		userRepository: userRepository,
		jwtService:     jwtService,
	}
}

func (s *AuthService) Login(
	ctx context.Context,
	login string,
	password string,
) (*repository.User, string, error) {
	user, err := s.userRepository.FindUser(ctx, login)
	if err != nil {
		return nil, "", err
	}

	err = s.checkPassword(user.PasswordHash, password)
	if err != nil {
		return nil, "", errors.New("invalid credentials")
	}

	role, err := s.userRepository.FindRoleByID(ctx, user.Role.ID)
	if err != nil {
		return nil, "", err
	}

	user.Role = *role

	token, err := s.jwtService.Generate(
		user.ID,
		user.Login,
		user.Role.Name,
	)
	if err != nil {
		return nil, "", err
	}

	return user, token, nil
}

func (s *AuthService) checkPassword(
	passwordHash string,
	password string,
) error {
	return bcrypt.CompareHashAndPassword(
		[]byte(passwordHash),
		[]byte(password),
	)
}
