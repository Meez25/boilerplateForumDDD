package http

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/meez25/boilerplateForumDDD/application/services"
	"github.com/meez25/boilerplateForumDDD/infrastructure/http/handlers"
	"github.com/meez25/boilerplateForumDDD/infrastructure/http/utils"
	"github.com/meez25/boilerplateForumDDD/infrastructure/persistence"
)

func StartServer() {
	r := chi.NewRouter()

	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	r.Use(middleware.Compress(5, "text/html", "text/css"))

	// Static files
	cfs := utils.CustomFileSystem{Fs: http.Dir("./ui")}
	fs := http.FileServer(cfs)

	// Intiate repo and services
	authenticationRepository := persistence.NewSessionMemoryRepo()
	authenticationService := services.NewAuthenticationService(authenticationRepository)

	// Route handlers
	r.Get("/", handlers.NewHomeHandler(*authenticationService).ServeHTTP)
	r.Get("/connexion/", handlers.NewLoginPageHandler(*authenticationService).ServeHTTP)
	r.Get("/inscription/", handlers.NewRegisterPageHandler(*authenticationService).ServeHTTP)
	r.Post("/connexion/", handlers.NewLoginPageHandler(*authenticationService).ServeHTTP)
	r.Get("/static/*", http.StripPrefix("/static", fs).ServeHTTP)

	// Start server
	http.ListenAndServe(":3000", r)
}
