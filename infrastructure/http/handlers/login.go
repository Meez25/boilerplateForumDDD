package handlers

import (
	"fmt"
	"net/http"

	"github.com/meez25/boilerplateForumDDD/application/services"
)

type LoginHandler struct {
	sessionServer services.AuthenticationService
}

func NewLoginHandler(sessionServer services.AuthenticationService) *LoginHandler {
	return &LoginHandler{sessionServer: sessionServer}
}

func (h *LoginHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {

	session, err := h.sessionServer.CreateSession("example@example.com")
	http.SetCookie(w, &http.Cookie{
		Name:     "sessionID",
		Value:    session.ID.String(),
		Expires:  session.GetValidUntil(),
		Secure:   true,
		HttpOnly: true,
	})
	fmt.Fprintf(w, "Your session is %v and error is %v", session, err)
}
