package handlers

import (
	"net/http"

	"github.com/meez25/boilerplateForumDDD/application/services"
	"github.com/meez25/boilerplateForumDDD/infrastructure/http/templates/auth"
	"github.com/meez25/boilerplateForumDDD/infrastructure/persistence"
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
	email := r.FormValue("email")
	password := r.FormValue("password")

	errors := make(map[string]string)

	user, err := h.sessionServer.Authenticate(email, password)

	if err != nil {
		switch err {
		case services.ErrInvalidCredentials:
			errors["general"] = "Le mot de passe est incorrect."
		case persistence.ErrUserAlreadyExists:
			errors["general"] = "L'utilisateur existe déjà."
		case persistence.ErrDuplicateEmail:
			errors["email"] = "L'adresse email est déjà utilisé"
		case persistence.ErrDuplicateUsername:
			errors["username"] = "Le nom d'utilisateur est déjà utilisé"
		default:
			errors["general"] = "Il y a eu une erreur"
		}

		auth.LoginForm(errors).Render(r.Context(), w)
		return
	}

	session, err := h.sessionServer.CreateSession(user.Email)

	http.SetCookie(w, &http.Cookie{
		Name:     "sessionID",
		Value:    session.ID.String(),
		Expires:  session.GetValidUntil(),
		Secure:   true,
		HttpOnly: true,
	})

	w.Header().Add("HX-Redirect", "/")
	http.Redirect(w, r, "/", http.StatusOK)
}
