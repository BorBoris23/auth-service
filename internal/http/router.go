package http

import (
	"net/http"
)

func NewRouter(authController *AuthController) http.Handler {
	mux := http.NewServeMux()

	mux.HandleFunc("POST /login", authController.Login)
	mux.HandleFunc("POST /register", authController.Register)

	return mux
}
