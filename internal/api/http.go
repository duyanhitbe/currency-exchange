package api

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
)

func GET(url string) map[string]interface{} {
	// Create a GET request
	resp, err := http.Get(url)
	if err != nil {
		log.Fatalf("Failed to make the request: %v", err)
	}
	defer resp.Body.Close()

	// Check the response status code
	if resp.StatusCode != http.StatusOK {
		log.Fatalf("Failed with status: %d", resp.StatusCode)
	}

	// Read the response body
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatalf("Failed to read the response body: %v", err)
	}

	var result map[string]interface{}

	// Unmarshal JSON data into the map
	err = json.Unmarshal(body, &result)
	if err != nil {
		log.Fatalf("Failed to unmarshal response body: %v", err)
	}

	return result
}
