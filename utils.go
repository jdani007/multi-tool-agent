package main

import (
	"context"
	"fmt"
	"os"

	"google.golang.org/adk/v2/agent"
	"google.golang.org/adk/v2/cmd/launcher"
	"google.golang.org/adk/v2/cmd/launcher/full"
	"google.golang.org/adk/v2/tool"
	"google.golang.org/adk/v2/tool/functiontool"
)

// MODEL is the Gemini model used by the agent for inference.
const MODEL = "gemini-3.1-flash-lite"

// newTool wraps a function tool with the metadata needed by the agent framework.
func newTool[TArgs, TResults any](name, description string, handler functiontool.Func[TArgs, TResults]) (tool.Tool, error) {
	tool, err := functiontool.New(
		functiontool.Config{
			Name:        name,
			Description: description,
		},
		handler,
	)
	if err != nil {
		return nil, err
	}

	return tool, nil
}

// launchAgent starts the interactive agent runner with the provided agent instance.
func launchAgent(ctx context.Context, a agent.Agent) error {
	launcherConfig := &launcher.Config{
		AgentLoader: agent.NewSingleLoader(a),
	}

	l := full.NewLauncher()
	err := l.Execute(ctx, launcherConfig, os.Args[1:])
	if err != nil {
		return fmt.Errorf("%v\n\n%s", err, l.CommandLineSyntax())
	}

	return nil
}
