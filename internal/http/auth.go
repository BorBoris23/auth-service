package http

import (
	"encoding/json"
	"net/http"

	"auth-service/internal/auth"
	"auth-service/internal/http/dto"

	"github.com/go-playground/validator/v10"
)

type AuthController struct {
	validate    *validator.Validate
	authService *auth.AuthService
}

func NewAuthController(
	validate *validator.Validate,
	authService *auth.AuthService,
) *AuthController {
	return &AuthController{
		validate:    validate,
		authService: authService,
	}
}

func (c *AuthController) Login(w http.ResponseWriter, r *http.Request) {
	var request dto.LoginRequest

	if !c.decodeAndValidate(
		w,
		r,
		&request,
		"Login and password are required",
	) {
		return
	}

	user, jwtToken, err := c.authService.Login(
		r.Context(),
		request.Login,
		request.Password,
	)
	if err != nil {
		writeError(
			w,
			http.StatusUnauthorized,
			"invalid_credentials",
			"Invalid login or password",
		)
		return
	}

	if err := writeJSON(
		w,
		http.StatusOK,
		newAuthResponse(user, jwtToken, "login successful"),
	); err != nil {
		http.Error(w, "internal server error", http.StatusInternalServerError)
	}
}

func (c *AuthController) Register(w http.ResponseWriter, r *http.Request) {
	var request dto.RegisterRequest

	if !c.decodeAndValidate(
		w,
		r,
		&request,
		"Name, login and password are required",
	) {
		return
	}

	user, jwtToken, err := c.authService.Register(
		r.Context(),
		request.Name,
		request.Login,
		request.Password,
	)
	if err != nil {
		writeError(
			w,
			http.StatusConflict,
			"registration_failed",
			"User with this login already exists",
		)
		return
	}

	if err := writeJSON(
		w,
		http.StatusCreated,
		newAuthResponse(user, jwtToken, "registration successful"),
	); err != nil {
		http.Error(w, "internal server error", http.StatusInternalServerError)
	}
}

func (c *AuthController) decodeAndValidate(
	w http.ResponseWriter,
	r *http.Request,
	request any,
	message string,
) bool {
	err := json.NewDecoder(r.Body).Decode(request)
	if err != nil {
		writeError(
			w,
			http.StatusBadRequest,
			"invalid_request",
			"Invalid request body",
		)
		return false
	}

	err = c.validate.Struct(request)
	if err != nil {
		writeError(
			w,
			http.StatusBadRequest,
			"validation_failed",
			message,
		)
		return false
	}

	return true
}
