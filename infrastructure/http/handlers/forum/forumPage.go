package forum

import (
	"fmt"
	"net/http"

	"github.com/meez25/boilerplateForumDDD/application/services"
	"github.com/meez25/boilerplateForumDDD/infrastructure/http/templates/forum"
	"github.com/meez25/boilerplateForumDDD/internal/authentication"
)

type ForumPageHandler struct {
	categoryService services.CategoryService
	userService     services.UserService
}

func NewForumPageHandler(categoryService services.CategoryService, userService services.UserService) *ForumPageHandler {
	return &ForumPageHandler{
		categoryService: categoryService,
		userService:     userService,
	}
}

func (fp *ForumPageHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	categories, err := fp.categoryService.GetAllCategoryAndChildren()

	if err != nil {
		fmt.Println("Could not retrieve categories")
	}

	context := r.Context().Value("session").(authentication.Session)
	user, err := fp.userService.FindByEmailAddress(context.Email)

	if err != nil {
		forum.Forum(categories, nil).Render(r.Context(), w)
		return
	}

	forum.Forum(categories, &user).Render(r.Context(), w)
}
