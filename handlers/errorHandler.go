package handlers

import (
	"net/http"

	"github.com/kwamekyeimonies/Rss-Aggregator/helper"
)

func HandlerErr(w http.ResponseWriter, r *http.Request) {
	helper.RespondWithError(w, 400, "Something went wrong")
}
