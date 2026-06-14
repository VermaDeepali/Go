package provider

import (
	"bytes"
	"encoding/json"
	"errors"
	"net/http"
	"os"
)

func CallOpenAI(prompt string) (string, error) {

	apiKey := os.Getenv("OPENROUTER_API_KEY")

	if apiKey == "" {
		return "", errors.New("OPENROUTER_API_KEY is missing")
	}

	body := map[string]any{
		"model": "openai/gpt-4o-mini",
		"messages": []map[string]string{
			{
				"role":    "user",
				"content": prompt,
			},
		},
	}

	jsonBody, err := json.Marshal(body)
	if err != nil {
		return "", err
	}

	req, err := http.NewRequest(
		"POST",
		"https://openrouter.ai/api/v1/chat/completions",
		bytes.NewBuffer(jsonBody),
	)

	if err != nil {
		return "", err
	}

	req.Header.Set("Authorization", "Bearer "+apiKey)
	req.Header.Set("Content-Type", "application/json")

	// Optional but recommended by OpenRouter
	req.Header.Set("HTTP-Referer", "http://localhost:8080")
	req.Header.Set("X-Title", "AI Gateway Go")

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

	choices, ok := result["choices"].([]any)

	if !ok || len(choices) == 0 {
		return "", errors.New("OpenRouter API returned invalid response")
	}

	message, ok := choices[0].(map[string]any)["message"].(map[string]any)

	if !ok {
		return "", errors.New("message parsing failed")
	}

	content, ok := message["content"].(string)

	if !ok {
		return "", errors.New("content parsing failed")
	}

	return content, nil
}
