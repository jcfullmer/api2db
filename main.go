package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/jcfullmer/api2db/internal/database"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

type Config struct {
	ApiKey  string
	dbQuery *database.Queries
}

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file: %s", err)
		os.Exit(2)
	}
	apiKey := os.Getenv("API_KEY")
	dbURL := os.Getenv("DB_URL")

	db, err := sql.Open("postgres", dbURL)
	if err != nil {
		log.Fatalf("Error connecting to DB: %v", err)
	}
	dbQuery := database.New(db)
	config := Config{
		ApiKey:  apiKey,
		dbQuery: dbQuery,
	}
	limit := 50
	start := 0
	for start <= 500 {
		fmt.Printf("New request starting at entry %v", start)
		requestURL := fmt.Sprintf("https://developer.nps.gov/api/v1/parks?limit=%d&start=%d", limit, start)
		err = config.RequestParseAdd(requestURL)
		if err != nil {
			log.Fatalf("%v\n", err)
		}
		start += 50
	}
}

func (conf Config) RequestParseAdd(url string) error {
	client := &http.Client{}
	req, err := http.NewRequest(http.MethodGet, url, nil)
	req.Header.Add("X-Api-Key", conf.ApiKey)
	req.Header.Add("accept", "application/json")
	if err != nil {
		log.Fatalf("Error creating request: %s", err)
		os.Exit(2)
	}
	resp, err := client.Do(req)
	if err != nil {
		log.Fatalf("Error getting response: %s", err)
		os.Exit(2)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		log.Printf("Bad status code: %v\n", resp.StatusCode)
		os.Exit(2)
	}

	err = ResponseToDB(resp, conf.dbQuery)
	if err != nil {
		log.Fatalf("error putting response into database: %v", err)
	}
	return nil
}
