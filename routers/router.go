package routers

import (
	"net/http"

	"github.com/FadyGamilM/Rss-Aggregator/handlers"
	"github.com/go-chi/chi"
	"github.com/go-chi/cors"
)

func Routes() http.Handler {
	router := chi.NewRouter()
	router.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"http://*", "https://*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "PATCH", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: true,
		MaxAge:           300,
	}))

	health_router := chi.NewRouter()
	health_router.Get("/health", handlers.Test_server_health)
	health_router.Get("/error", handlers.Test_error_handler)

	router.Mount("/api/v1", health_router)

	return router
}
