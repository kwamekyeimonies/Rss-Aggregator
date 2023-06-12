package apirouter

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/kwamekyeimonies/Rss-Aggregator/handlers"
)

func ApiRouter(r chi.Router) {
	r.Get("/ready", handlers.HandlerRequest)
	r.Get("/error", handlers.HandlerErr)
	r.Post("/user", func(w http.ResponseWriter, r *http.Request) {
		apiConfig := &handlers.NewApiConfig{}
		apiConfig.HandleCreateuser(w, r)
	})
	r.Get("/user", func(w http.ResponseWriter, r *http.Request) {
		apiConfig := &handlers.NewApiConfig{}
		apiConfig.HandleGetUsers(w, r)
	})
}
