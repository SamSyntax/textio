package internal

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/cors"
	ready "github.com/samsyntax/textio/internal/handlers"
)

func InitRoutes() http.Handler {
	router := chi.NewRouter()

	router.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"https://*", "http://*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"*"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: false,
		MaxAge:           300,
	}))

	// v1 router 
	v1Router := chi.NewRouter()
  // Routes
	v1Router.Get("/healthz", ready.HandlerReadiness)
  v1Router.Get("/err", ready.HandlerErr)
  // Mounting to the /v1 route
	router.Mount("/v1", v1Router)

	return router

}
