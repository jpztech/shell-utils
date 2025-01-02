package ai

import (
	"log"
	"strings"

	"shell-utils/ai/gemini"
	"shell-utils/ai/model"
	"shell-utils/rest"
)

func QA(args []string) string {
	query := strings.Join(args, " ")
	request, err := client.LLM.Request(query)
	if err != nil {
		log.Fatalf("Failed to create request: %v", err)
	}
	response, err := client.RestClient.Do(request)
	if err != nil {
		log.Fatalf("Failed to do request: %v", err)
	}
	result, err := client.LLM.FromResponse(response)
	if err != nil {
		log.Fatalf("Failed to parse response: %v", err)
	}
	return result
}

type AIClient struct {
	model.LLM
	RestClient *rest.Client
}

var client *AIClient

func InitAIClient(config *Config) {
	models := make(map[string]model.LLM)
	models["gemini"] = gemini.NewLLM(config.URL, config.Properties)
	if model, ok := models[config.Model]; ok {
		client = &AIClient{
			LLM: model,
			RestClient: rest.NewClient(config.Proxy),
		}
	} else {
		log.Fatalf("Model %s not found", config.Model)
	}
}

func GetClient() *AIClient {
	return client
}