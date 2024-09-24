package handlers

import (
	"net/http"

	"github.com/meez25/boilerplateForumDDD/application/services"
)

type LogoutHandler struct {
	authService services.AuthenticationService
}

func NewLogoutHandler(AuthService services.AuthenticationService) *LogoutHandler {
	return &LogoutHandler{
		authService: AuthService,
	}
}

func (h *LogoutHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	sessionCookie, err := r.Cookie("sessionID")

	if err != nil {
		http.Redirect(w, r, "/", http.StatusTemporaryRedirect)
	}

	err = h.authService.Logout(sessionCookie.Value)

	if err != nil {
		http.Redirect(w, r, "/", http.StatusTemporaryRedirect)
	}

}
