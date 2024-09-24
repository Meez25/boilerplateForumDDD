package handlers

import (
	"fmt"
	"net/http"

	"github.com/meez25/boilerplateForumDDD/application/services"
)

type LoginPageHandler struct {
	sessionServer services.AuthenticationService
}

func NewLoginPageHandler(sessionServer services.AuthenticationService) *LoginPageHandler {
	return &LoginPageHandler{sessionServer: sessionServer}
}

func (h *LoginPageHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	session, err := h.sessionServer.CreateSession("example@example.com")
	fmt.Fprintf(w, "Hello, World! Your session is %v and error is %v", session, err)
}
