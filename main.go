package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/joho/godotenv"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

func main() {
	apiURL := "https://api.mistral.ai/v1/chat/completions"
	err := godotenv.Load()
	if err != nil {
		log.Println("No .env file found")
	}
	apiKey := os.Getenv("MISTRAL_API_KEY")
	if apiKey == "" {
		log.Fatal("MISTRAL_API_KEY not set")
	}

	payload := map[string]interface{}{
		"model": "mistral-small-latest",
		"messages": []map[string]interface{}{
			{
				"content": "Tell me a joke about Go programming.",
				"role":    "user",
			},
		},
	}
	body, _ := json.Marshal(payload)

	req, _ := http.NewRequest("POST", apiURL, bytes.NewBuffer(body))
	req.Header.Set("Authorization", "Bearer "+apiKey)
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	respBody, _ := ioutil.ReadAll(resp.Body)
	fmt.Println("Mistral response:", string(respBody))
}
