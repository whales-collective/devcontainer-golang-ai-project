package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/firebase/genkit/go/ai"
	"github.com/firebase/genkit/go/genkit"
	"github.com/firebase/genkit/go/plugins/compat_oai/openai"
	"github.com/openai/openai-go/option"
)

func main() {
	ctx := context.Background()

	modelRunnerBaseUrl := os.Getenv("MODEL_RUNNER_BASE_URL")
	modelRunnerChatModel := os.Getenv("MODEL_RUNNER_CHAT_MODEL")

	g := genkit.Init(ctx, genkit.WithPlugins(&openai.OpenAI{
		APIKey: "I💙DockerModelRunner",
		Opts: []option.RequestOption{
			option.WithBaseURL(modelRunnerBaseUrl),
		},
	}))

	_, err := genkit.Generate(ctx, g,
		ai.WithModelName("openai/"+modelRunnerChatModel),
		//ai.WithModelName("ai/qwen2.5:0.5B-F16"),

		ai.WithMessages(
			ai.NewSystemTextMessage("You are the dungeon master of a D&D game."),
			ai.NewUserTextMessage("Generate a D&D NPC Elf name and all its characteristics."),
		),
		ai.WithConfig(map[string]any{"temperature": 0.7}),

		ai.WithStreaming(func(ctx context.Context, chunk *ai.ModelResponseChunk) error {
			// Do something with the chunk...
			fmt.Print(chunk.Text())
			return nil
		}),
	)
	if err != nil {
		log.Fatal(err)
	}

}
