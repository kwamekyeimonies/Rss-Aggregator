package apirouter

import (
	"github.com/go-chi/chi/v5"
	"github.com/kwamekyeimonies/Rss-Aggregator/service"
)

func ApiRouter(r chi.Router) {
	r.Get("/ready", service.HandlerRequest)
	r.Get("/error", service.HandlerErr)
}
