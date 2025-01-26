package utils

import (
	"io"
	"log"
	"os"
)

type SystemPrompt struct {
	Prompt string
}

func ReadFile(fileUrl string) (*SystemPrompt, error) {
	file, err := os.Open(fileUrl)
	if err != nil {
		return nil, err

	}
	defer func() {
		if err = file.Close(); err != nil {
			log.Fatal(err)
		}
	}()

	b, err := io.ReadAll(file)

	prompt := &SystemPrompt{
		Prompt: string(b),
	}
	return prompt, nil
}
