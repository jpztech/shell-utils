package ai

import (
	"fmt"
	"log"

	"shell-utils/ai/gemini"
	"shell-utils/ai/model"
	"shell-utils/rest"
	"shell-utils/viewer"
)

func DoQA(query string) string {
	return callLLM(query, model.QA).(string)
}

func DoShellAssistant(query string) model.ShellAssistantResponse {
	return callLLM(query, model.ShellAssistant).(model.ShellAssistantResponse)
}

func callLLM(query string, scenario model.Scenario) interface{} {
	request, err := client.LLM.Request(query, scenario)
	if err != nil {
		log.Fatalf("Failed to create request: %v", err)
	}
	bar := viewer.Loading("Waiting response from LLM...")
	response, err := client.RestClient.Do(request)
	if err != nil {
		bar.Finish()
		fmt.Printf("\r\033[K")
		log.Fatalf("Failed to do request: %v", err)
	}
	bar.Finish()
	fmt.Printf("\r\033[K")
	result, err := client.LLM.DecodeResponse(response, scenario)
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