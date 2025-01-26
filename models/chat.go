package models

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
	Stream            bool                     `json:"stream,omitempty"`
	Response_Format   Format                   `json:"response_format"`
}
