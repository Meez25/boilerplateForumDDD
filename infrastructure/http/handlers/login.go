package handlers

import (
	"fmt"
	"net/http"

	"github.com/meez25/boilerplateForumDDD/application/services"
	"github.com/meez25/boilerplateForumDDD/infrastructure/http/templates/auth"
)

type LoginHandler struct {
	sessionServer services.AuthenticationService
}

func NewLoginHandler(sessionServer services.AuthenticationService) *LoginHandler {
	return &LoginHandler{
		sessionServer: sessionServer,
	}
}

func (h *LoginHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Login handler")
	email := r.FormValue("email")
	password := r.FormValue("password")

	fmt.Println(email, password)

	errors := make(map[string]string)

	user, err := h.sessionServer.Authenticate(email, password)

	if err != nil {
		switch err {
		case services.ErrInvalidCredentials:
			errors["general"] = "Le mot de passe est incorrect."
		default:
			errors["general"] = "Il y a eu une erreur"
		}
	}

	fmt.Println(user)

	session, err := h.sessionServer.CreateSession(user.Email)
	http.SetCookie(w, &http.Cookie{
		Name:     "sessionID",
		Value:    session.ID.String(),
		Expires:  session.GetValidUntil(),
		Secure:   true,
		HttpOnly: true,
	})

	auth.LoginForm(errors).Render(r.Context(), w)
}
