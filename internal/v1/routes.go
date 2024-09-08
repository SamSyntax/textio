package internal

import (
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/cors"
	"github.com/samsyntax/textio/internal/handlers"
	"net/http"
)

func InitRoutes(apiCfg handlers.ApiConfig) http.Handler {
	router := chi.NewRouter()

	router.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"https://*", "http://*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Content-Type", "Authorization"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: false,
		MaxAge:           301,
	}))

	// v1 router
	v1Router := chi.NewRouter()
	// Routes
	v1Router.Get("/healthz", handlers.HandlerReadiness)
	v1Router.Get("/err", handlers.HandlerErr)
	v1Router.Post("/users", apiCfg.HandlerCreateUser)
	v1Router.Get("/users", apiCfg.MiddlewareAuth(apiCfg.HandlerGetUser))

	// feed routes
	v1Router.Post("/feeds", apiCfg.MiddlewareAuth(apiCfg.HandlerCreateFeed))
	v1Router.Get("/feeds", apiCfg.HandlerGetFeeds)

	//feed follows
	v1Router.Post("/feed_follows", apiCfg.MiddlewareAuth(apiCfg.HandlerCreateFeedFollow))
	v1Router.Get("/feed_follows", apiCfg.MiddlewareAuth(apiCfg.HandlerGetFeedFollow))
	v1Router.Delete("/feed_follows/{feedFollowID}", apiCfg.MiddlewareAuth(apiCfg.HandlerDeleteFeedFollow))

  // Posts routes
  v1Router.Get("/posts", apiCfg.MiddlewareAuth(apiCfg.HandlerGetPostsForUser))


  // Mounting to the /v1 route
	router.Mount("/v1", v1Router)

	return router

}
