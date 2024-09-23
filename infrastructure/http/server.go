package http

import (
	"net/http"

	"github.com/meez25/boilerplateForumDDD/application/services"
	"github.com/meez25/boilerplateForumDDD/infrastructure/http/handlers"
	"github.com/meez25/boilerplateForumDDD/infrastructure/persistence"
)

func StartServer() {
	authenticationRepository := persistence.NewSessionMemoryRepo()
	authenticationService := services.NewAuthenticationService(authenticationRepository)
	http.HandleFunc("/", handlers.NewHomeHandler(*authenticationService).ServeHTTP)

	http.ListenAndServe(":3000", nil)
}
