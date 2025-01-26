package utils

import (
	"bufio"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"

	"github.com/lew-monk/biashara.git/models"
)

type ChatCompletionStreamInterface interface {
	Recv() (*models.StreamChatCompletionResponse, error)
	Close() error
}

type ChatCompletionStream struct {
	Ctx    context.Context
	Cancel context.CancelFunc
	Resp   *http.Response
	Reader *bufio.Reader
}

func (s *ChatCompletionStream) Recv() (*models.StreamChatCompletionResponse, error) {

	reader := s.Reader

	// Infinite Loop to read the incoming string data
	for {
		// Read the current message until we get a new line
		line, err := reader.ReadString('\n')

		if err != nil {
			// Check if we have gotten to the end off line and return EOF error
			if err == io.EOF {
				return nil, io.EOF
			}
			return nil, fmt.Errorf("error reading the output stream: %w", err)
		}

		line = strings.TrimSpace(line)

		// Check if the stream is done
		if line == "data: [DONE]" {
			return nil, io.EOF
		}

		// Check if the stream has any text and the text is part of the data
		if len(line) > 6 && line[:6] == "data: " {
			// Remove the data prefix
			cleanText := line[6:]

			fmt.Println(cleanText)

			// Declare a variable to store the incoming response
			var streamResponse models.StreamChatCompletionResponse

			// Unmarshall the response to our struct
			if err := json.Unmarshal([]byte(cleanText), &streamResponse); err != nil {
				return nil, fmt.Errorf(
					"Couldnt deciphy cleaned text from: %s, error is %w",
					cleanText,
					err,
				)
			}

			// if its the first chunk set the stream usage to default
			if streamResponse.Usage == nil {
				streamResponse.Usage = &models.StreamUsage{}
			}

			return &streamResponse, nil
		}
	}

}

func (s *ChatCompletionStream) CloseStream() error {
	s.Cancel()
	return s.Resp.Body.Close()
}
