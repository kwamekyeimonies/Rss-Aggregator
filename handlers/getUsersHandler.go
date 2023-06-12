package handlers

import (
	"fmt"
	"net/http"

	"github.com/kwamekyeimonies/Rss-Aggregator/helper"
	"github.com/kwamekyeimonies/Rss-Aggregator/internal/auth"
	"github.com/kwamekyeimonies/Rss-Aggregator/models"
)

func (ApiCfg *NewApiConfig) HandleGetUsers(w http.ResponseWriter, r *http.Request) {
	apiKey, err := auth.GetApiKey(r.Header)
	if err != nil {
		helper.RespondWithError(w, 403, fmt.Sprintf("Auth error: %v\n", err.Error()))
		return
	}

	user, err := ApiCfg.PostgreSQL_DB_API().DB.GetUserByAPIKey(r.Context(), apiKey)
	if err != nil {
		helper.RespondWithError(w, 403, fmt.Sprintf("Auth Error,Reason: %v\n", err.Error()))
		return
	}

	helper.RespondWithJson(w, 200, models.DatabaseuserToUser(user))
}
