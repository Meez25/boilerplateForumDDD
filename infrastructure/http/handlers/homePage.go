package handlers

import (
	"net/http"

	"github.com/meez25/boilerplateForumDDD/application/services"
	"github.com/meez25/boilerplateForumDDD/infrastructure/http/templates"
	"github.com/meez25/boilerplateForumDDD/internal/authentication"
)

type HomeHandler struct {
	sessionService services.AuthenticationService
}

func NewHomeHandler(sessionService services.AuthenticationService) *HomeHandler {
	return &HomeHandler{
		sessionService: sessionService,
	}
}

func (h *HomeHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	context := r.Context().Value("session").(authentication.Session)

	templates.Index(context).Render(r.Context(), w)
}
