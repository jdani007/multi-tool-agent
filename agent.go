package main

import (
	"context"
	"log"
	"os"

	"google.golang.org/adk/v2/agent/llmagent"
	"google.golang.org/adk/v2/model/gemini"
	"google.golang.org/adk/v2/tool"
	"google.golang.org/genai"

	"multi-agent-tool/tools"
	"multi-agent-tool/helpers"

	_ "time/tzdata"
)

func main() {
	ctx := context.Background()

	model, err := gemini.NewModel(ctx, helpers.MODEL, &genai.ClientConfig{
		APIKey: os.Getenv("GOOGLE_API_KEY"),
	})
	if err != nil {
		log.Fatalf("Failed to create model: %v", err)
	}

	weatherTool, err := helpers.NewTool("getWeather", "Retrieves the current weather report for a specified city.", tools.GetWeather)
	if err != nil {
		log.Fatalf("Failed to create getWeather tool: %v", err)
	}

	currentTimeTool, err := helpers.NewTool("getCurrentTime", "Returns the current time in a specified city.", tools.GetCurrentTime)
	if err != nil {
		log.Fatalf("Failed to create getCurrentTime tool: %v", err)
	}

	rootAgent, err := llmagent.New(llmagent.Config{
		Name:        "Weather and Time Agent",
		Model:       model,
		Description: "Agent to answer questions about the time and weather in a city.",
		Instruction: `You are a helpful agent who can answer user questions about the time and weather in a city.
						All metrics are returned in American Standard units (Fahrenheit, mph, etc.) by default. Provide conversions if requested.
						Use all of the data returned from the tools to provide the most appropriate responses.
						You can also have a conversation on the data that you report. One question or statement at a time.
						`,
		Tools: []tool.Tool{
			weatherTool,
			currentTimeTool,
		},
	})

	if err != nil {
		log.Fatalf("Failed to create agent: %v", err)
	}

	if err := helpers.LaunchAgent(ctx, rootAgent); err != nil {
		log.Fatalf("Run failed: %v", err)
	}
}
