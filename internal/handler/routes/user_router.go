package routes

import (
	"github.com/go-chi/chi"
	"github.com/nathanfabio/completeAPIGo/internal/handler/userhandler"
)

func InitUserRoutes(router chi.Router, h userhandler.UserHandler) {
	router.Route("/user", func(r chi.Router) {
		r.Post("/", nil)
	})
}