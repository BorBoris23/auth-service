package grpc

import (
	"context"

	"auth-service/internal/jwt"

	authpb "github.com/BorBoris23/auth-proto/gen/auth"
)

type AuthServer struct {
	authpb.UnimplementedAuthServiceServer

	jwtService *jwt.JWTService
}

func NewAuthServer(jwtService *jwt.JWTService) *AuthServer {
	return &AuthServer{
		jwtService: jwtService,
	}
}

func (s *AuthServer) ValidateToken(
	ctx context.Context,
	req *authpb.ValidateTokenRequest,
) (*authpb.ValidateTokenResponse, error) {
	token, err := s.jwtService.Validate(req.Token)
	if err != nil {
		return &authpb.ValidateTokenResponse{
			Valid: false,
		}, nil
	}

	claims, ok := token.Claims.(*jwt.Claims)
	if !ok {
		return &authpb.ValidateTokenResponse{
			Valid: false,
		}, nil
	}

	return &authpb.ValidateTokenResponse{
		Valid:  true,
		UserId: int64(claims.UserID),
	}, nil
}
