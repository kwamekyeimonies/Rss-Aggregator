package handlers

import (
	"fmt"
	"net/http"

	"github.com/kwamekyeimonies/Rss-Aggregator/helper"
	"github.com/kwamekyeimonies/Rss-Aggregator/models"
)

func (pgCfg *NewApiConfig) GetAllFeeds(w http.ResponseWriter, r *http.Request) {
	feeds, err := pgCfg.PostgreSQL_DB_API().DB.GetFeeds(r.Context())
	if err != nil {
		helper.RespondWithError(w, 400, fmt.Sprintf("No feeds in db %v\n", err.Error()))
		return
	}

	helper.RespondWithJson(w, 200, models.DatabaseFeedstoFeeds(feeds))
}
