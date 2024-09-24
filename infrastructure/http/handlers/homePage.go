package handlers

import (
	"fmt"
	"net/http"

	"github.com/meez25/boilerplateForumDDD/application/services"
	"github.com/meez25/boilerplateForumDDD/infrastructure/http/templates"
)

type HomeHandler struct {
	sessionService services.AuthenticationService
}

func NewHomeHandler(sessionServer services.AuthenticationService) *HomeHandler {
	return &HomeHandler{sessionService: sessionServer}
}

func (h *HomeHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	cookie, err := r.Cookie("sessionID")

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(cookie)
		_, err := h.sessionService.GetSessionByID(cookie.Value)
		if err != nil {
			fmt.Println("La session n'existe plus")
		}
	}

	// session, err := h.sessionServer.CreateSession("example@example.com")
	// http.SetCookie(w, &http.Cookie{
	// 	Name:     "sessionID",
	// 	Value:    session.ID.String(),
	// 	Expires:  session.GetValidUntil(),
	// 	Secure:   true,
	// 	HttpOnly: true,
	// })
	// fmt.Fprintf(w, "Hello, World! Your session is %v and error is %v", session, err)
	templates.Index().Render(r.Context(), w)
}
