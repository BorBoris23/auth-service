package handlers

import (
	"encoding/json"
	"net/http"

	"auth-service/internal/server/dto"
	"auth-service/internal/server/service"

	"github.com/go-playground/validator/v10"
)

type AuthHandler struct {
	validate    *validator.Validate
	authService *service.AuthService
}

func NewAuthHandler(
	validate *validator.Validate,
	authService *service.AuthService,
) *AuthHandler {
	return &AuthHandler{
		validate:    validate,
		authService: authService,
	}
}

func (h *AuthHandler) Login(w http.ResponseWriter, r *http.Request) {
	var request dto.LoginRequest

	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		writeJSON(w, http.StatusBadRequest, dto.ErrorResponse{
			Error:   "invalid_request",
			Message: "Invalid request body",
		})
		return
	}

	err = h.validate.Struct(request)
	if err != nil {
		writeJSON(w, http.StatusBadRequest, dto.ErrorResponse{
			Error:   "validation_failed",
			Message: "Login and password are required",
		})
		return
	}

	user, jwtToken, err := h.authService.Login(
		r.Context(),
		request.Login,
		request.Password,
	)
	if err != nil {
		writeJSON(w, http.StatusUnauthorized, dto.ErrorResponse{
			Error:   "invalid_credentials",
			Message: "Invalid login or password",
		})
		return
	}

	response := dto.LoginResponse{
		Message: "login successful",
		Token:   jwtToken,
		User: dto.UserResponse{
			ID:    user.ID,
			Login: user.Login,
			Role:  user.Role.Name,
		},
	}

	if err := writeJSON(w, http.StatusOK, response); err != nil {
		http.Error(w, "internal server error", http.StatusInternalServerError)
	}
}
