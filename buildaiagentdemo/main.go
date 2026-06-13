package main

import (
	"bufio"
	"context"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/fatih/color"
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

	// messages = append(messages, openai.UserMessage(("what model are you??")))

	params := openai.ChatCompletionNewParams{
		Model:    model,
		Messages: messages,
	}

	scanner := bufio.NewScanner(os.Stdin)

	fmt.Println(color.WhiteString(
		fmt.Sprintf("model: %s", model),
	))

	for {
		fmt.Println(color.CyanString("\n> "))

		if !scanner.Scan() {
			break
		}

		input := strings.TrimSpace(scanner.Text())

		if input == "clear" {
			fmt.Print("\033[H\033[]2J")
			continue
		}

		if input == "exit" {
			fmt.Println("goodbye")
			break
		}

		params.Messages = append(params.Messages, openai.UserMessage(input))

		res, err := client.Chat.Completions.New(ctx, params)

		if err != nil {
			log.Fatal(err)
		}
		output := res.Choices[0].Message.Content

		fmt.Println(color.YellowString(output))

		params.Messages = append(params.Messages, openai.AssistantMessage(output))

		count := len(params.Messages)

		fmt.Println(color.MagentaString(
			fmt.Sprintf("count: %d", count),
		))

		if count >= 10 {
			params.Messages = params.Messages[count-4:]
			fmt.Println("history refreshed!")
		}
	}

}
