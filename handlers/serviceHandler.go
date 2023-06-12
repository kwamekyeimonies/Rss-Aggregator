package handlers

import (
	"net/http"

	"github.com/kwamekyeimonies/Rss-Aggregator/helper"
)

func HandlerRequest(w http.ResponseWriter, r *http.Request) {
	helper.RespondWithJson(w, http.StatusOK, struct{}{})
}
