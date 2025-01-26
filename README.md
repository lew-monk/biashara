# DeepSeek API Usage

This guide provides instructions for using the DeepSeek API package effectively. Follow the steps below to set up and make requests to the API.

---

## Required Environment Variables

To use the DeepSeek API package, ensure the following environment variables are set:

1. `DEEP_SEEK_API_KEY`
2. `DEEP_SEEK_COMPLETIONS_URL`
3. `DEEP_SEEK_SYSTEM_PROMPT_TEXT_URL` (Optional)

> **Note:** Environment variables should be used instead of hardcoding values to enhance security and flexibility.

---

## Quick Start Guide

### Creating a Basic DeepSeek Request

To create a chat completion request:

```go

	resp, err := handlers.CreateChatCompletion(
		context,
		payload,
		DEEP_SEEK_API_KEY,
		DEEP_SEEK_COMPLETIONS_URL,
	)

	if err != nil {
		log.Fatal(err)
	}
```

**Enhance Assistant Output with System Prompts**

To refine the assistant's responses, you can leverage a system prompt. This is a set of instructions that guide the model's behavior and output.

- **How it works:**
  - Create a text file named `DEEP_SEEK_SYSTEM_PROMPT_TEXT_URL`.
  - Populate this file with your desired system-level instructions.
  - Pass this file as an input to the assistant's request.

**Example Instructions (in `DEEP_SEEK_SYSTEM_PROMPT_TEXT_URL`):**

- **Concise and Factual:** Respond with brief, informative answers.
- **Avoid Jargon:** Use plain language that is easy to understand.
- **Cite Sources:** When appropriate, provide links to reliable sources.
- **Maintain Neutrality:** Present information objectively without bias.

By incorporating a system prompt, you can tailor the assistant's responses to your specific needs and preferences.

Example:

```
	systemPrompt, err := utils.ReadFile(DEEP_SEEK_SYSTEM_PROMPT_TEXT_URL)
	system_desc := models.DeepSeekRequestMessage{}
```

## Using System Prompts with the DEEP_SEEK_SYSTEM_PROMPT_TEXT_URL Environment Variable

The `DEEP_SEEK_SYSTEM_PROMPT_TEXT_URL` environment variable specifies the relative path to a text file containing system-level instructions for guiding the model's behavior. This file is located in the root directory of your project.

Full example for a chat completion request:

```go

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

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

	resp, err := handlers.CreateChatCompletion(
		context,
		payload,
		DEEP_SEEK_API_KEY,
		DEEP_SEEK_COMPLETIONS_URL,
	)

	if err != nil {
		log.Fatal(err)
	}

    llmResp, err := utils.HandleChatResponse(resp)


	if err != nil {
		log.Fatal(err)
	}

    fmt.Println(llmResp)

```

## Streaming Responses

If you want to stream the response, there is a utility function that does the same.
You have to set `stream` to `true` on the payload for this to work.
It is important to note that the setup till sending the request is the same.

```go

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

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
```

How you handle the data after recieving the stream is upto you.
