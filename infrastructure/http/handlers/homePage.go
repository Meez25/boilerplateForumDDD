package handlers

import (
	"fmt"
	"net/http"

	"github.com/meez25/boilerplateForumDDD/application/services"
	"github.com/meez25/boilerplateForumDDD/infrastructure/http/templates"
)

type HomeHandler struct {
	bh             BaseHandler
	sessionService services.AuthenticationService
}

func NewHomeHandler(sessionService services.AuthenticationService) *HomeHandler {
	return &HomeHandler{
		bh:             *NewBaseHandler(sessionService),
		sessionService: sessionService,
	}
}

func (h *HomeHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	context := r.Context().Value("session")
	fmt.Println("context :", context)
	session, _ := h.bh.GetSession(w, r)

	fmt.Println("Session", session)

	templates.Index(session).Render(r.Context(), w)
}
