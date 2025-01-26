package utils

import (
	"fmt"
	"net/http"
	"time"
)

func HandleSendRequest(req *http.Request) (*http.Response, error) {
	client := &http.Client{
		Timeout: 240 * time.Second,
	}

	res, err := client.Do(req)

	if err != nil {
		panic(
			fmt.Sprintf("failed to send request: %v", err),
		)
	}

	return res, nil
}
