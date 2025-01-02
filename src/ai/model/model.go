package model

import "net/http"

type LLM interface {
	Request(query string) (*http.Request, error)
	FromResponse(response *http.Response) (string, error)
}
