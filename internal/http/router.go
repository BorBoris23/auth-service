package http

import (
	"net/http"
)

func NewRouter(authController *AuthController) http.Handler {
	mux := http.NewServeMux()

	mux.HandleFunc("/login", authController.Login)
	mux.HandleFunc("/register", authController.Register)

	return mux
}
