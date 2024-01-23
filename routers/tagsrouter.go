package routes

import (
	"Portfoliobackend/services"
	"github.com/go-chi/chi/v5"
	"net/http"
)

func TagsRoutes() http.Handler {
	router := chi.NewRouter()
	router.Get("/", services.GetTags)

	return router
}
