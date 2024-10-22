package admin

import (
	"fmt"
	"net/http"

	"github.com/meez25/boilerplateForumDDD/application/services"
	"github.com/meez25/boilerplateForumDDD/infrastructure/http/templates/admin"
	"github.com/meez25/boilerplateForumDDD/internal/authentication"
)

type AdminCategoryPageHandler struct {
	categoryService services.CategoryService
	userService     services.UserService
}

func NewAdminCategoryPageHandler(categoryService services.CategoryService, userService services.UserService) *AdminCategoryPageHandler {
	return &AdminCategoryPageHandler{
		categoryService: categoryService,
		userService:     userService,
	}
}

func (ap *AdminCategoryPageHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
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

	categories, err := ap.categoryService.GetAllCategoryAndChildren()

	if err != nil {
		fmt.Println("Could not retrieve categories")
	}

	admin.AdminCategory(categories).Render(r.Context(), w)
}
