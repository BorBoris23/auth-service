package http

import (
	"auth-service/internal/http/dto"
	"auth-service/internal/user"
	"encoding/json"
	"net/http"
)

func writeJSON(w http.ResponseWriter, status int, response any) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)

	return json.NewEncoder(w).Encode(response)
}

func writeError(
	w http.ResponseWriter,
	status int,
	errorCode string,
	message string,
) {
	writeJSON(w, status, dto.ErrorResponse{
		Error:   errorCode,
		Message: message,
	})
}

func newLoginResponse(user *user.User, jwtToken string) dto.LoginResponse {
	return dto.LoginResponse{
		Message: "login successful",
		Token:   jwtToken,
		User: dto.UserResponse{
			ID:    user.ID,
			Login: user.Login,
			Role:  user.Role.Name,
		},
	}
}
