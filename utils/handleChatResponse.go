package utils

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/lew-monk/biashara/models"
)

func HandleResponseChat(resp *http.Response) (*models.ChatCompletionResponse, error) {
	body, err := io.ReadAll(resp.Body)

	if err != nil {
		panic(
			fmt.Sprintf("failed to read body: %v", err),
		) // Using panic for debugging; replace in future with proper error handling.
	}

	defer resp.Body.Close()

	fmt.Println(string(body))

	var responeBody models.ChatCompletionResponse

	if err := json.Unmarshal(body, &responeBody); err != nil {
		return nil, fmt.Errorf("failed to parse response JSON: %w", err)
	}

	return &responeBody, nil

}
