package handlers

import (
	"bufio"
	"context"
	"errors"
	"log"

	"github.com/lew-monk/biashara.git/models"
	"github.com/lew-monk/biashara.git/utils"
)

func CreateStreamChatCompletion(
	ctx context.Context,
	payload *models.DeepSeekRequest,
	DEEP_SEEK_API_KEY string,
	DEEP_SEEK_COMPLETIONS_URL string,
) (*utils.ChatCompletionStream, error) {

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

	if res.StatusCode >= 400 {
		return &utils.ChatCompletionStream{}, errors.New("Failed")
	}

	ctx, cancel := context.WithCancel(ctx)

	stream := &utils.ChatCompletionStream{
		Ctx:    ctx,
		Cancel: cancel,
		Resp:   res,
		Reader: bufio.NewReader(res.Body),
	}

	return stream, nil
}
