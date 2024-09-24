package handlers

import (
	"net/http"

	"github.com/meez25/boilerplateForumDDD/application/services"
	"github.com/meez25/boilerplateForumDDD/internal/authentication"
)

type BaseHandler struct {
	Fs                    http.HandlerFunc
	authenticationService services.AuthenticationService
}

func NewBaseHandler(authService services.AuthenticationService) *BaseHandler {
	return &BaseHandler{
		authenticationService: authService,
	}
}

func (bh *BaseHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	bh.Fs.ServeHTTP(w, r)
}

func (bh BaseHandler) GetSession(w http.ResponseWriter, r *http.Request) (authentication.Session, error) {
	cookie, err := r.Cookie("sessionID")
	session := authentication.Session{}

	if err != nil {
		return session, err
	}

	session, err = bh.authenticationService.GetSessionByID(cookie.Value)

	if err != nil {
		return session, err
	}

	return session, nil
}
