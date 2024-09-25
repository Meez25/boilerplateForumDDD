package handlers

import (
	"fmt"
	"net/http"

	"github.com/meez25/boilerplateForumDDD/application/services"
	"github.com/meez25/boilerplateForumDDD/infrastructure/http/templates/auth"
	"github.com/meez25/boilerplateForumDDD/infrastructure/persistence"
	"github.com/meez25/boilerplateForumDDD/internal/user"
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

	errors := make(map[string]string)
	// values := make(map[string]string)
	values := map[string]string{
		"username":         username,
		"email":            email,
		"password":         password,
		"confirm-password": confirmPassword,
	}

	createdUser, err := h.userService.Create(username, email, password, confirmPassword, "firstName", "lastName", "PP")

	if err != nil {
		fmt.Println(err)
		switch err {
		case services.ErrPasswordConfirmError:
			errors["confirm-password"] = "Les mots de passe ne correspondent pas"
		case user.ErrEmptyEmail:
			errors["email"] = "L'adresse email ne peut pas être vide"
		case user.ErrEmptyUsername:
			errors["username"] = "Le nom d'utilisateur ne peut pas être vide"
		case user.ErrEmptyPassword:
			errors["username"] = "Le mot de passe ne peut pas être vide"
		case persistence.ErrEmailAlreadyExists:
			errors["email"] = "L'adresse email est déjà utilisée"
		case persistence.ErrUserAlreadyExists:
			errors["email"] = "L'utilisateur existe déjà"
		case persistence.ErrDuplicateUsername:
			errors["username"] = "Le nom d'utilisateur est déjà utilisé"
		case persistence.ErrDuplicateEmail:
			errors["email"] = "L'adresse email est déjà utilisée"
		default:
			errors["general"] = "Une erreur inattendue s'est produite. Veuillez réessayer."
		}
	} else {
		for k := range values {
			delete(values, k)
		}
		values["general"] = "Bienvenue " + createdUser.Username + " !"
	}

	auth.RegisterForm(errors, values).Render(r.Context(), w)
}
