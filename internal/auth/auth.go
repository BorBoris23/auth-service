package auth

import (
	"context"
	"errors"

	"auth-service/internal/jwt"
	"auth-service/internal/role"
	"auth-service/internal/user"

	"golang.org/x/crypto/bcrypt"
)

type AuthService struct {
	userRepository *user.UserRepository
	roleRepository *role.RoleRepository
	jwtService     *jwt.JWTService
}

func NewAuthService(
	userRepository *user.UserRepository,
	roleRepository *role.RoleRepository,
	jwtService *jwt.JWTService,
) *AuthService {
	return &AuthService{
		userRepository: userRepository,
		roleRepository: roleRepository,
		jwtService:     jwtService,
	}
}

func (s *AuthService) Login(
	ctx context.Context,
	login string,
	password string,
) (*user.User, string, error) {
	user, err := s.userRepository.FindUser(ctx, login)
	if err != nil {
		return nil, "", err
	}

	err = s.checkPassword(user.PasswordHash, password)
	if err != nil {
		return nil, "", errors.New("invalid credentials")
	}

	role, err := s.roleRepository.FindByID(ctx, user.Role.ID)
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
