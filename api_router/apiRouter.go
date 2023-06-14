package apirouter

import (
	"github.com/go-chi/chi/v5"
	dbconnection "github.com/kwamekyeimonies/Rss-Aggregator/dbConnection"
	"github.com/kwamekyeimonies/Rss-Aggregator/handlers"
	"github.com/kwamekyeimonies/Rss-Aggregator/middleware"
)

func ApiRouter(r chi.Router) {

	apiConfig := &handlers.NewApiConfig{
		NewApiConfig: &dbconnection.NewApiConfig{},
	}

	mwConfig := &middleware.NewApiConfig{
		NewApiConfig: &dbconnection.NewApiConfig{},
	}

	r.Get("/ready", handlers.HandlerRequest)
	r.Get("/error", handlers.HandlerErr)
	r.Post("/user", apiConfig.HandleCreateuser)
	r.Get("/user", mwConfig.MiddlewareAuth(apiConfig.HandleGetUsers))
	r.Post("/feed", mwConfig.MiddlewareAuth(apiConfig.HandlerCreateUserFeed))
	r.Get("/feed", apiConfig.GetAllFeeds)
}
