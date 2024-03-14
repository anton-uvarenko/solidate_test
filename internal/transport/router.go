package transport

import (
	"github.com/go-chi/chi/v5"
	"net/http"
	"solidgate_test/internal/transport/handlers"
)

func NewRouter(handler *handlers.Handler) *http.Server {
	mux := chi.NewRouter()
	setUpRoutes(mux, handler)

	s := &http.Server{
		Handler: mux,
		Addr:    ":8080",
	}
	return s
}

func setUpRoutes(router *chi.Mux, handler *handlers.Handler) {
	router.Post("/validate", handler.CardValidatorHandler.IsCardValid)
}
