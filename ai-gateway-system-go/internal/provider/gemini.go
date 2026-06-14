package provider

import (
	"bytes"
	"encoding/json"
	"errors"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
)

func CallGemini(prompt string) (string, error) {
	// Load env file
	err := godotenv.Load()

	if err != nil {
		log.Fatal("Error loading .env file")
	}

	apiKey := os.Getenv("GEMINI_API_KEY")

	if apiKey == "" {
		return "", errors.New("GEMINI_API_KEY is missing")
	}

	body := map[string]any{
		"contents": []map[string]any{
			{
				"parts": []map[string]string{
					{
						"text": prompt,
					},
				},
			},
		},
	}

	jsonBody, err := json.Marshal(body)
	if err != nil {
		return "", err
	}

	url := "https://generativelanguage.googleapis.com/v1beta/models/gemini-flash-latest:generateContent?key=" + apiKey

	req, err := http.NewRequest(
		"POST",
		url,
		bytes.NewBuffer(jsonBody),
	)

	if err != nil {
		return "", err
	}

	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}

	resp, err := client.Do(req)
	if err != nil {
		return "", err
	}

	defer resp.Body.Close()

	var result map[string]any

	err = json.NewDecoder(resp.Body).Decode(&result)
	if err != nil {
		return "", err
	}

	// DEBUG RESPONSE
	pretty, _ := json.MarshalIndent(result, "", "  ")
	println(string(pretty))

	candidates, ok := result["candidates"].([]any)

	if !ok || len(candidates) == 0 {
		return "", errors.New("Gemini API returned invalid response")
	}

	contentMap := candidates[0].(map[string]any)["content"].(map[string]any)

	parts := contentMap["parts"].([]any)

	text := parts[0].(map[string]any)["text"].(string)

	return text, nil
}
