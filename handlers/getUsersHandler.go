package handlers

import (
	"net/http"

	"github.com/kwamekyeimonies/Rss-Aggregator/helper"
	"github.com/kwamekyeimonies/Rss-Aggregator/internal/database"
	"github.com/kwamekyeimonies/Rss-Aggregator/models"
)

func (ApiCfg *NewApiConfig) HandleGetUsers(w http.ResponseWriter, r *http.Request, user database.User) {

	helper.RespondWithJson(w, 200, models.DatabaseuserToUser(user))
}
