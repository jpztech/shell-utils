package gemini

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"

	"shell-utils/ai/model"
)

type Gemini struct {
	URL 		string
	Properties 	map[string]string
}

type Part struct {
	Model 		string `json:"model,omitempty"`
	Text 		string `json:"text"`
}

type Content struct {
	Role 		string `json:"role,omitempty"`
	Parts 		[]Part `json:"parts"`
}

type RequestBody struct {
	Contents 			[]Content `json:"contents"`
	GenerationConfig 	map[string]interface{} `json:"generationConfig,omitempty"`
}

type Candidate struct {
	Content 		Content `json:"content"`
	FinishReason 	string `json:"finishReason"`
}

type ResponseBody struct {
	Candidates 	[]Candidate `json:"candidates"`
	ModelVersion 	string `json:"modelVersion"`
}

const SHELL_ASSISTANT_GENERATION_CONFIG string = `{
    "responseMimeType": "application/json",
    "responseSchema": {
        "type": "object",
        "properties": {
            "command": {
                "type": "string",
                "description": "The shell command for completing the task."
            },
            "explanation": {
                "type": "string",
                "description": "A brief explanation of what the shell command does."
            }
        },
        "required": [
            "command",
            "explanation"
        ]
    }
}`

func (g *Gemini) Request(query string, scenario model.Scenario) (*http.Request, error) {
	body := RequestBody{
		Contents: []Content{
			{
				Parts: []Part{
					{
						Text: query,
					},
				},
			},
		},
	}
	if scenario == model.ShellAssistant {
		body.GenerationConfig = map[string]interface{}{}
		err := json.Unmarshal([]byte(SHELL_ASSISTANT_GENERATION_CONFIG), &body.GenerationConfig)
		if err != nil {
			return nil, err
		}
	}
	jsonBody, err := json.Marshal(body)
	if err != nil {
		return nil, err
	}
	r, err := http.NewRequest("POST", g.URL, bytes.NewBuffer(jsonBody))
	if err != nil {
		return nil, err
	}
	r.Header.Set("Content-Type", "application/json")
	return r, nil
}

func (g *Gemini) DecodeResponse(response *http.Response, scenario model.Scenario) (interface{}, error) {
	body := ResponseBody{}
	defer response.Body.Close()
	err := json.NewDecoder(response.Body).Decode(&body)
	if err != nil {
		return "", err
	}
	if len(body.Candidates) == 0 || len(body.Candidates[0].Content.Parts) == 0 {
		body, _ := json.MarshalIndent(body, "", "  ")
		return "", fmt.Errorf("invalid response: %s", string(body))
	}
	text := body.Candidates[0].Content.Parts[0].Text
	if scenario == model.ShellAssistant {
		result := model.ShellAssistantResponse{}
		err = json.Unmarshal([]byte(text), &result)
		if err != nil {
			fmt.Println("Failed to unmarshal response")
			return "", err
		}
		return result, nil
	}
	return text, nil
}

func NewLLM(url string, properties map[string]string) model.LLM {
	return &Gemini{
		URL: url,
		Properties: properties,
	}
}