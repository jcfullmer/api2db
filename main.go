package main

import (
	"database/sql"
	"log"
	"net/http"
	"os"

	"github.com/jcfullmer/api2db/internal/database"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

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

	client := &http.Client{}
	// requestURL := "https://developer.nps.gov/api/v1/newsreleases?limit=5"
	requestURL := "https://developer.nps.gov/api/v1/parks?stateCode=ID&limit=1"

	req, err := http.NewRequest(http.MethodGet, requestURL, nil)
	req.Header.Add("X-Api-Key", apiKey)
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

	err = ResponseToDB(resp, dbQuery)
	if err != nil {
		log.Fatalf("error putting response into database: %v", err)
	}
}
