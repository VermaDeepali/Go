package model

type AskRequest struct {
	Prompt string `json:"prompt"`
	Model  string `json:"model"`
}
