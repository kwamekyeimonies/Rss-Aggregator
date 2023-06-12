package main

import (
	"log"
	"net/http"
	"os"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/cors"
	"github.com/joho/godotenv"
	apirouter "github.com/kwamekyeimonies/Rss-Aggregator/api_router"
)

func main() {

	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal(err)
	}

	serverPort := os.Getenv("PORT")
	if serverPort == "" {
		log.Fatal("PORT is not found in the Environment")
	}

	router := chi.NewRouter()

	router.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"https://*", "http://*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: false,
		MaxAge:           300, // Maximum value not ignored by any of maj
	}))

	router.Route("/api", apirouter.ApiRouter)

	server := http.Server{
		Handler: router,
		Addr:    ":" + serverPort,
	}
	log.Printf("Server running on port: %v\n", serverPort)

	err = server.ListenAndServe()
	if err != nil {
		log.Fatal(err.Error())
	}

}
