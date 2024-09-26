package admin

import (
	"fmt"
	"net/http"

	"github.com/meez25/boilerplateForumDDD/application/services"
	"github.com/meez25/boilerplateForumDDD/infrastructure/http/templates/admin"
	"github.com/meez25/boilerplateForumDDD/internal/authentication"
)

type AdminCategoryHandler struct {
	categoryService services.CategoryService
	userService     services.UserService
}

func NewAdminCategoryHandler(categoryService services.CategoryService, userService services.UserService) *AdminCategoryHandler {
	return &AdminCategoryHandler{
		categoryService: categoryService,
		userService:     userService,
	}
}

func (ap *AdminCategoryHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
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

	title := r.FormValue("title")
	description := r.FormValue("description")

	c, err := ap.categoryService.Create(title, description, nil)

	if err != nil {
		fmt.Println("Could not create the category")
	}

	category := services.CategoryAndChildren{
		Category: *c,
	}

	admin.CategoryItem(category).Render(r.Context(), w)
}
