package handlers

import (
	"fmt"
	"net/http"

	"github.com/meez25/boilerplateForumDDD/application/services"
)

// HomeHandler is a simple handler

type HomeHandler struct {
	sessionServer services.AuthenticationService
}

func NewHomeHandler(sessionServer services.AuthenticationService) *HomeHandler {
	return &HomeHandler{sessionServer: sessionServer}
}

func (h *HomeHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	session, err := h.sessionServer.CreateSession("example@example.com")
	fmt.Fprintf(w, "Hello, World! Your session is %v and error is %v", session, err)
}
