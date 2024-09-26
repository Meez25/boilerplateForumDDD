package admin

import (
	"net/http"

	"github.com/meez25/boilerplateForumDDD/application/services"
	"github.com/meez25/boilerplateForumDDD/infrastructure/http/templates/admin"
	"github.com/meez25/boilerplateForumDDD/internal/authentication"
)

type AdminPageHandler struct {
	categoryService services.CategoryService
	userService     services.UserService
}

func NewAdminPageHandler(categoryService services.CategoryService, userService services.UserService) *AdminPageHandler {
	return &AdminPageHandler{
		categoryService: categoryService,
		userService:     userService,
	}
}

func (ap *AdminPageHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	// Are you even an admin ?!
	context := r.Context().Value("session").(authentication.Session)
	user, err := ap.userService.FindByEmailAddress(context.Email)

	if err != nil {
		http.Redirect(w, r, "/connexion", http.StatusTemporaryRedirect)
		return
	}

	if !user.SuperAdmin {
		http.Redirect(w, r, "/", http.StatusTemporaryRedirect)
		return
	}

	admin.AdminIndex().Render(r.Context(), w)
}
