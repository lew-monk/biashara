package examples

import (
	"context"
	"errors"
	"fmt"
	"io"
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/lew-monk/biashara/constants"
	"github.com/lew-monk/biashara/handlers"
	"github.com/lew-monk/biashara/models"
	"github.com/lew-monk/biashara/utils"
)

type Environment struct {
	// Environment variables
	DEEP_SEEK_API_KEY string `env:"DEEP_SEEK_API_KEY, required=true"`
}

func main() {

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	// This bit is user option. How you pass the required params is upto you. That is the api key as well as the deepseek url

	DEEP_SEEK_API_KEY := os.Getenv("DEEP_SEEK_API_KEY")
	DEEP_SEEK_COMPLETIONS_URL := os.Getenv("DEEP_SEEK_COMPLETIONS_URL")
	DEEP_SEEK_SYSTEM_PROMPT_TEXT_URL := os.Getenv("DEEP_SEEK_SYSTEM_PROMPT_TEXT_URL")

	systemPrompt, err := utils.ReadFile(DEEP_SEEK_SYSTEM_PROMPT_TEXT_URL)
	system_desc := models.DeepSeekRequestMessage{}

	if err != nil {
		fmt.Errorf(
			"System Prompt Not Provided. The software still works but it will be more efficient with more information on what to produce",
		)
	} else {

		system_desc = models.DeepSeekRequestMessage{
			Content: systemPrompt.Prompt,
			Role:    constants.ChatMessageRoleSystem,
		}
	}

	// Needed to make the request
	payload := &models.DeepSeekRequest{
		Model:             constants.ChatModel,
		Frequency_penalty: 0,
		Stream:            true,
		Response_Format: models.Format{
			Type: "text",
		},
	}

	if systemPrompt != nil {
		payload.Messages = []models.DeepSeekRequestMessage{
			system_desc,
			models.DeepSeekRequestMessage{
				Content: "Create a shoe ecommerce website. I mainly focus on Nikes and Addidas. Kindly include images of the latest release and have different shoe sizes on the product. The theme color should be black and white",
				Role:    "user",
			},
		}

	} else {
		payload.Messages = []models.DeepSeekRequestMessage{
			models.DeepSeekRequestMessage{
				Content: "Create a shoe ecommerce website. I mainly focus on Nikes and Addidas. Kindly include images of the latest release and have different shoe sizes on the product. The theme color should be black and white. Dont use a boring gray. Let it be exciting you can have black but not too much to avoid contrash issues",
				Role:    "user",
			},
		}
	}

	context := context.Background()

	// This is a normal chat completion without streaming.

	/* resp, err := handlers.CreateChatCompletion(
		context,
		payload,
		DEEP_SEEK_API_KEY,
		DEEP_SEEK_COMPLETIONS_URL,
	)

	if err != nil {
		log.Fatal(err)
	} */

	// THis with streaming implemented. You can create a go routine that sends the information to the fontend as the stream message is updated. Whatever you do with the stream data is upto you

	stream, err := handlers.CreateStreamChatCompletion(
		context,
		payload,
		DEEP_SEEK_API_KEY,
		DEEP_SEEK_COMPLETIONS_URL,
	)

	if err != nil {
		log.Fatalf("ChatCompletionStream error: %v", err)
	}

	var fullMessage string

	defer stream.CloseStream()

	for {
		recv, err := stream.Recv()

		if errors.Is(err, io.EOF) {
			fmt.Println("\nStream finished")
			break
		}

		if err != nil {
			fmt.Printf("\nStream error: %v\n", err)
			break
		}

		for _, choice := range recv.Choices {
			fullMessage += choice.Delta.Content // Accumulate chunk content
			log.Println(fullMessage)
		}
	}

	fmt.Println(fullMessage)

}
