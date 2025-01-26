package handlers

import (
	"context"
	"log"

	"github.com/lew-monk/biashara.git/models"
	"github.com/lew-monk/biashara.git/utils"
)

func CreateChatCompletion(
	ctx context.Context,
	payload *models.DeepSeekRequest,
	DEEP_SEEK_API_KEY string,
	DEEP_SEEK_COMPLETIONS_URL string,
) (*models.ChatCompletionResponse, error) {

	client, err := utils.NewRequest(DEEP_SEEK_API_KEY)

	if err != nil {
		log.Fatal(err)
	}

	req, err := client.SetBaseUrl(DEEP_SEEK_COMPLETIONS_URL).
		SetBodyFromStruct(payload).
		Build(ctx)

	res, err := utils.HandleSendRequest(req)

	if err != nil {
		log.Fatal(err)
	}

	resp, err := utils.HandleResponseChat(res)

	if err != nil {
		log.Fatal(err)
	}

	return resp, nil
}
