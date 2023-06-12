package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/google/uuid"
	dbconnection "github.com/kwamekyeimonies/Rss-Aggregator/dbConnection"
	"github.com/kwamekyeimonies/Rss-Aggregator/helper"
	"github.com/kwamekyeimonies/Rss-Aggregator/internal/database"
	"github.com/kwamekyeimonies/Rss-Aggregator/models"
)

type NewApiConfig struct {
	*dbconnection.NewApiConfig
}

func (ApiCfg *NewApiConfig) HandleCreateuser(w http.ResponseWriter, r *http.Request) {
	type parameters struct {
		Name string `json:"name"`
	}

	decoder := json.NewDecoder(r.Body)
	params := parameters{}

	err := decoder.Decode(&params)

	if err != nil {
		helper.RespondWithError(w, 400, fmt.Sprintf("Error passing JSON,Reason: %v\n", err.Error()))
		return
	}

	user, err := ApiCfg.PostgreSQL_DB_API().DB.CreateUser(r.Context(), database.CreateUserParams{
		ID:        uuid.New(),
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
		Name:      params.Name,
	})

	if err != nil {
		helper.RespondWithError(w, 400, fmt.Sprintf("Couldn't create user: %v\n", err.Error()))
	}

	helper.RespondWithJson(w, 200, models.DatabaseuserToUser(user))
}
