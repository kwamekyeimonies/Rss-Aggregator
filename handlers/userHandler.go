package handlers

import (
	"net/http"

	"github.com/kwamekyeimonies/Rss-Aggregator/models"
)

type NewApiConfig struct {
	*models.ApiConfig
}

func (ApiCfg *NewApiConfig) HandleCreateuser(w http.ResponseWriter, r *http.Request) {

}
