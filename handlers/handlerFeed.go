package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/google/uuid"
	"github.com/kwamekyeimonies/Rss-Aggregator/helper"
	"github.com/kwamekyeimonies/Rss-Aggregator/internal/database"
	"github.com/kwamekyeimonies/Rss-Aggregator/models"
)

func (apiCfg *NewApiConfig) HandlerCreateUserFeed(w http.ResponseWriter, r *http.Request, user database.User) {
	type parameters struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	}

	decoder := json.NewDecoder(r.Body)
	params := parameters{}

	err := decoder.Decode(&params)
	if err != nil {
		helper.RespondWithError(w, 400, fmt.Sprintf("Error passing JSON: %v\n", err.Error()))
		return
	}

	fmt.Println(params)

	userFeed, err := apiCfg.PostgreSQL_DB_API().DB.CreatedFeed(r.Context(), database.CreatedFeedParams{
		ID:        uuid.New(),
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
		Name:      params.Name,
		Url:       params.URL,
		UserID:    user.ID,
	})

	if err != nil {
		helper.RespondWithError(w, 400, fmt.Sprintf("Couldn't create user: %v\n", err.Error()))
		return
	}

	helper.RespondWithJson(w, 200, models.DatabaseFeedtoFeed(userFeed))
}
