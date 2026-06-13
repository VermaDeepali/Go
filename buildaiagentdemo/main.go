package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/openai/openai-go/v2"
	"github.com/openai/openai-go/v2/option"
)

func main() {

	// Load env file
	err := godotenv.Load()

	if err != nil {
		log.Fatal("Error loading .env file")
	}

	apiKey := os.Getenv("OPENROUTER_API_KEY")

	// log.Println("API key loaded", apiKey)

	if apiKey == "" {
		log.Fatal("api key is required!")
	}

	url := "https://openrouter.ai/api/v1"

	client := openai.NewClient(
		option.WithBaseURL(url),
		option.WithAPIKey(apiKey),
	)

	ctx := context.Background()

	messages := []openai.ChatCompletionMessageParamUnion{}

	// model := "nvidia/nemotron-3-ultra-550b-a55b:free"
	// model := "google/gemini-2.0-flash-001"
	model := "openai/gpt-4o-mini"

	messages = append(messages, openai.UserMessage(("what model are you??")))

	params := openai.ChatCompletionNewParams{
		Model:    model,
		Messages: messages,
	}

	res, err := client.Chat.Completions.New(ctx, params)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(res.Choices[0].Message.Content)

}
