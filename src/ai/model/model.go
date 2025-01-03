package model

import "net/http"

type LLM interface {
	Request(query string, scenario Scenario) (*http.Request, error)
	DecodeResponse(response *http.Response, scenario Scenario) (interface{}, error)
}

// define an enum for scenarios
type Scenario int
const (
	QA Scenario = iota
	ShellAssistant
)

type ShellAssistantResponse struct {
	Command 	string `json:"command"`
	Explanation string `json:"explanation"`
}