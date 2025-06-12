package ai

import (
	"context"
	"fmt"
	"github.com/openai/openai-go"
	"github.com/openai/openai-go/option"
	"os"
)

var apiKey = os.Getenv("API_KEY")
var apiModel = os.Getenv("API_MODEL")

func Tts(request string) {
	if len(apiModel) == 0 {
		fmt.Println("You must provide a command for help with ai")
		return
	}
	if len(apiKey) == 0 {
		fmt.Println("please provide env variable API_KEY")
		return
	}
	client := openai.NewClient(option.WithAPIKey(apiKey))
	stream := client.Chat.Completions.NewStreaming(context.TODO(), openai.ChatCompletionNewParams{
		Messages: []openai.ChatCompletionMessageParamUnion{
			openai.UserMessage(request),
		},
		Model: apiModel,
	})

	for stream.Next() {
		chunk := stream.Current()
		fmt.Print(chunk.Choices[0].Delta.Content)
	}

}
