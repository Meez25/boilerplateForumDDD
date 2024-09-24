package middleware

import (
	"context"
	"net/http"

	"github.com/meez25/boilerplateForumDDD/application/services"
	"github.com/meez25/boilerplateForumDDD/internal/authentication"
)

type AuthMiddlewareService struct {
	authService services.AuthenticationService
}

func NewAuthMiddlewareService(authService services.AuthenticationService) *AuthMiddlewareService {
	return &AuthMiddlewareService{
		authService: authService,
	}
}

func (au *AuthMiddlewareService) GetSessionInContext(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		session := authentication.Session{}

		cookie, err := r.Cookie("sessionID")

		if err != nil {
			ctx := context.WithValue(r.Context(), "session", session)
			next.ServeHTTP(w, r.WithContext(ctx))
			return
		}

		session, err = au.authService.GetSessionByID(cookie.Value)

		if err != nil {
			ctx := context.WithValue(r.Context(), "session", session)
			next.ServeHTTP(w, r.WithContext(ctx))
			return
		}

		ctx := context.WithValue(r.Context(), "session", session)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
