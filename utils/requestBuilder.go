package utils

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
)

type Format struct {
	Type string `json:"type"`
}

type DeepSeekRequestMessage struct {
	Content string `json:"content"`
	Role    string `json:"role"`
}

type DeepSeekRequest struct {
	Messages          []DeepSeekRequestMessage `json:"messages"`
	Model             string                   `json:"model"`
	Frequency_penalty int                      `json:"frequency_penalty"`
	Stream            bool                     `json:"stream"`
	Response_Format   Format                   `json:"response_format"`
}

type AuthenticatedRequest struct {
	BearerToken string
	BaseUrl     string
	Path        string
	Body        []byte
}

func (rb *AuthenticatedRequest) SetBaseUrl(url string) *AuthenticatedRequest {
	rb.BaseUrl = url
	return rb
}

func (rb *AuthenticatedRequest) SetPath(path string) *AuthenticatedRequest {
	rb.Path = path
	return rb
}

func (rb *AuthenticatedRequest) SetBodyFromStruct(
	body interface{},
) *AuthenticatedRequest {

	payload, err := json.Marshal(body)
	if err != nil {
		panic(
			fmt.Sprintf("failed to marshal body: %v", err),
		) // Using panic for debugging; replace in future with proper error handling.
	}

	rb.Body = payload

	return rb
}

func NewRequest(token string) (*AuthenticatedRequest, error) {
	if token == "" {
		return nil, errors.New("API Key Must be provided")
	}

	bearer := fmt.Sprintf("Bearer %s", token)

	rb := &AuthenticatedRequest{
		BearerToken: bearer,
	}

	return rb, nil
}

func (rb *AuthenticatedRequest) Build(ctx context.Context) (*http.Request, error) {

	if rb.BaseUrl == "" || rb.BearerToken == "" {
		return nil, errors.New("Base Url not set")
	}

	req, err := http.NewRequestWithContext(
		ctx,
		http.MethodPost,
		rb.BaseUrl,
		bytes.NewReader(rb.Body),
	)

	if err != nil {
		return nil, err
	}

	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Accept", "application/json")
	req.Header.Add("Authorization", rb.BearerToken)

	return req, nil

}
