package handlers

import (
	"fmt"
	"net/http"

	"github.com/meez25/boilerplateForumDDD/application/services"
)

type RegisterHandler struct {
	sessionService services.AuthenticationService
	userService    services.UserService
}

func NewRegisterHandler(sessionService services.AuthenticationService, userService services.UserService) *RegisterHandler {
	return &RegisterHandler{
		sessionService: sessionService,
		userService:    userService,
	}
}

func (h *RegisterHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	username := r.FormValue("username")
	email := r.FormValue("email")
	password := r.FormValue("password")
	confirmPassword := r.FormValue("confirm-password")

	user, err := h.userService.Create(username, email, password, confirmPassword, "firstName", "lastName", "PP")

	if err != nil {
		fmt.Fprintf(w, "erreur %v !", err)
		w.WriteHeader(http.StatusBadRequest)
	}

	session, err := h.sessionService.CreateSession(user.EmailAddress)
	http.SetCookie(w, &http.Cookie{
		Name:     "sessionID",
		Value:    session.ID.String(),
		Expires:  session.GetValidUntil(),
		Secure:   true,
		HttpOnly: true,
		Path:     "/",
	})
	fmt.Fprintf(w, "Hello, World! Your session is %v and error is %v", session, err)
}
