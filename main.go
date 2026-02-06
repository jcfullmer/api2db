package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file: %s", err)
		os.Exit(2)
	}
	apiKey := os.Getenv("API_KEY")
	client := &http.Client{}
	// requestURL := "https://developer.nps.gov/api/v1/newsreleases?limit=5"
	requestURL := "https://developer.nps.gov/api/v1/parks?stateCode=ID"

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

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Println(err)
		os.Exit(2)
	}
	fmt.Printf(string(body))
}
