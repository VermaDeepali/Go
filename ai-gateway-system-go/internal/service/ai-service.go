package service

import "ai-gateway-system-go/internal/provider"

func ProcessRequest(prompt string, model string) (string, error) {

	switch model {
	case "openai":
		return provider.CallOpenAI(prompt)
	case "gemini":
		return provider.CallGemini(prompt)
	default:
		return provider.CallOpenAI(prompt)
	}
}
