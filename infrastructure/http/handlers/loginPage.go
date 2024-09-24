package handlers

import (
	"net/http"

	"github.com/meez25/boilerplateForumDDD/application/services"
	"github.com/meez25/boilerplateForumDDD/infrastructure/http/templates/auth"
)

type LoginPageHandler struct {
	sessionServer services.AuthenticationService
}

func NewLoginPageHandler(sessionServer services.AuthenticationService) *LoginPageHandler {
	return &LoginPageHandler{sessionServer: sessionServer}
}

func (h *LoginPageHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	auth.Login(map[string]string{}).Render(r.Context(), w)
}
