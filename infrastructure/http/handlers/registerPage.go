package handlers

import (
	"net/http"

	"github.com/meez25/boilerplateForumDDD/application/services"
	"github.com/meez25/boilerplateForumDDD/infrastructure/http/templates/auth"
)

type RegisterPageHandler struct {
	sessionServer services.AuthenticationService
}

func NewRegisterPageHandler(sessionServer services.AuthenticationService) *RegisterPageHandler {
	return &RegisterPageHandler{sessionServer: sessionServer}
}

func (h *RegisterPageHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	auth.Register(map[string]string{}).Render(r.Context(), w)
}
