package tools

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

const MODEL = "gemini-3.1-flash-lite"

func NewTool[TArgs, TResults any](name, description string, handler functiontool.Func[TArgs, TResults]) (tool.Tool, error) {
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

func LaunchAgent(ctx context.Context, a agent.Agent) error {
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
