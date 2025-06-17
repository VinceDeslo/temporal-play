package internal

import (
	"context"
	"errors"
	"os"

	"google.golang.org/genai"
)

type LLMService struct {
	Model string
	Client *genai.Client 
}

func NewLLMService(ctx context.Context) (*LLMService, error) {
	apiKey := os.Getenv("GEMINI_API_KEY")
	if apiKey == "" {
		return nil, errors.New("gemini api key undefined")
	}

	clientCfg := &genai.ClientConfig{
		APIKey: apiKey,
	}
	client, err := genai.NewClient(ctx, clientCfg)
	if err != nil {
		return nil, errors.Join(errors.New("failed to generate genai client"), err)
	}

	return &LLMService{
		Model: "gemini-2.0-flash",
		Client: client,
	}, nil
}

func (llm *LLMService) Prompt(ctx context.Context, prompt string) (string, error){
	result, err := llm.Client.Models.GenerateContent(
		ctx,
		llm.Model,
		genai.Text(prompt),
		nil,
	)
	if err != nil {
		return "", errors.Join(errors.New("failed to prompt genai client"), err)
	}
	return result.Text(), nil
}
