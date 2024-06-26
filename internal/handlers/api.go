package handlers

import (
	"goapi/internal/handlers/middleware"

	"github.com/go-chi/chi"
	chimiddle "github.com/go-chi/chi/middleware"
)

func Handler(r *chi.Mux) {
	// Global middleware
	r.Use(chimiddle.StripSlashes)

	r.Get("/", Landing)

	r.Route("/account", func(router chi.Router) {

		// Middleware for /account route
		router.Use(middleware.Authorization)

		router.Get("/bitcoin", GetBitcoinBalance)
		router.Get("/ethereum", GetEthereumBalance)
	})
}
