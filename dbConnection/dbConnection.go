package dbconnection

import (
	"database/sql"
	"log"
	"os"

	"github.com/kwamekyeimonies/Rss-Aggregator/internal/database"
	"github.com/kwamekyeimonies/Rss-Aggregator/models"
	_ "github.com/lib/pq"
)

type NewApiConfig struct {
	*models.ApiConfig
}

func (Db_apiCfg *NewApiConfig) PostgreSQL_DB_API() *NewApiConfig {
	dbURL := os.Getenv("POSTGRES_DB_Connection")

	if dbURL == "" {
		log.Fatal("Postgres DB URL Not Found in the Environment")
	}

	dbConn, err := sql.Open("postgres", dbURL)
	if err != nil {
		log.Fatal("Cannot connect to Postgres Database")
	}

	queries := database.New(dbConn)
	if err != nil {
		log.Fatal("Cannot connect to database: ", err.Error())
	}

	apiCfg := &NewApiConfig{
		ApiConfig: &models.ApiConfig{
			DB: queries,
		},
	}

	return apiCfg
}
