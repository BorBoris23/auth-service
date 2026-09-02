package auth

import (
	"context"
	"errors"

	"auth-service/internal/jwt"
	"auth-service/internal/repository/role"
	"auth-service/internal/repository/user"

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

	err = s.loadUserRole(ctx, user)
	if err != nil {
		return nil, "", err
	}

	token, err := s.generateToken(user)
	if err != nil {
		return nil, "", err
	}

	return user, token, nil
}

func (s *AuthService) Register(
	ctx context.Context,
	name string,
	login string,
	password string,
) (*user.User, string, error) {
	passwordHash, err := bcrypt.GenerateFromPassword(
		[]byte(password),
		bcrypt.DefaultCost,
	)
	if err != nil {
		return nil, "", err
	}

	newUser, err := s.userRepository.CreateUser(
		ctx,
		name,
		login,
		string(passwordHash),
	)
	if err != nil {
		return nil, "", err
	}

	err = s.loadUserRole(ctx, newUser)
	if err != nil {
		return nil, "", err
	}

	token, err := s.generateToken(newUser)
	if err != nil {
		return nil, "", err
	}

	return newUser, token, nil
}

func (s *AuthService) loadUserRole(
	ctx context.Context,
	u *user.User,
) error {
	role, err := s.roleRepository.FindByID(ctx, u.Role.ID)
	if err != nil {
		return err
	}

	u.Role = *role

	return nil
}

func (s *AuthService) generateToken(
	u *user.User,
) (string, error) {
	return s.jwtService.Generate(
		u.ID,
		u.Login,
		u.Role.Name,
	)
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
