package http

import (
	"net/http"
)

func NewRouter(authController *AuthController) http.Handler {
	mux := http.NewServeMux()

	mux.HandleFunc("/login", authController.Login)

	return mux
}
