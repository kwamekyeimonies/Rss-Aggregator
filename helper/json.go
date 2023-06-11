package helper

import (
	"encoding/json"
	"log"
	"net/http"
)

type ErrResponse struct {
	Error string `json:"error"`
}

func RespondWithJson(w http.ResponseWriter, code int, payload interface{}) {
	data, err := json.Marshal(payload)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Printf("Failed to marshal JSON response: %v\n", payload)
		return
	}

	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(data)
}

func RespondWithError(w http.ResponseWriter, code int, msg string) {
	if code > 400 {
		log.Println("Responding with 5XX Error: ", msg)
	}

	RespondWithJson(w, code, ErrResponse{
		Error: msg,
	})

}
