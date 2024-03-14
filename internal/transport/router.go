package transport

import (
	"github.com/go-chi/chi/v5"
	"net/http"
)

func NewRouter() *http.Server {
	s := &http.Server{
		Handler: chi.NewRouter(),
		Addr:    ":80",
	}
	return s
}

func setUpRoutes(router *chi.Mux) {

}
