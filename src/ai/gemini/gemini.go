package gemini

import (
	"bytes"
	"encoding/json"
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
	Contents 	[]Content `json:"contents"`
}

type Candidate struct {
	Content 		Content `json:"content"`
	FinishReason 	string `json:"finishReason"`
}

type ResponseBody struct {
	Candidates 	[]Candidate `json:"candidates"`
	ModelVersion 	string `json:"modelVersion"`
}

func (g *Gemini) Request(query string) (*http.Request, error) {
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

func (g *Gemini) FromResponse(response *http.Response) (string, error) {
	body := ResponseBody{}
	defer response.Body.Close()
	err := json.NewDecoder(response.Body).Decode(&body)
	if err != nil {
		return "", err
	}
	return body.Candidates[0].Content.Parts[0].Text, nil
}

func NewLLM(url string, properties map[string]string) model.LLM {
	return &Gemini{
		URL: url,
		Properties: properties,
	}
}