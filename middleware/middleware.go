package middleware

import (
	"fmt"
	"net/http"

	dbconnection "github.com/kwamekyeimonies/Rss-Aggregator/dbConnection"
	"github.com/kwamekyeimonies/Rss-Aggregator/helper"
	"github.com/kwamekyeimonies/Rss-Aggregator/internal/auth"
	"github.com/kwamekyeimonies/Rss-Aggregator/internal/database"
)

type AuthHandler func(http.ResponseWriter, *http.Request, database.User)

type NewApiConfig struct {
	*dbconnection.NewApiConfig
}

func (Cfg *NewApiConfig) MiddlewareAuth(handler AuthHandler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		apiKey, err := auth.GetApiKey(r.Header)
		if err != nil {
			helper.RespondWithError(w, 403, fmt.Sprintf("Auth error: ", err.Error()))
			return
		}

		user, err := Cfg.PostgreSQL_DB_API().DB.GetUserByAPIKey(r.Context(), apiKey)
		if err != nil {
			helper.RespondWithError(w, 400, fmt.Sprintf("Auth error: ", err.Error()))
			return
		}

		handler(w, r, user)

	}
}
