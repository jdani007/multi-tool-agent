# Multi-tool Weather and Time Agent

This project is a small Go-based agent demo built with the Google ADK. It uses a Gemini model plus two tools:

- Weather tool: queries the public wttr.in API for the requested city and returns current weather details.
- Current time tool: returns the current time for an IANA time zone such as America/New_York or Europe/Paris.

The implementation is based on the [ADK multi-tool agent tutorial](https://adk.dev/tutorials/multi-tool-agent/), with the weather tool expanded to pull live data from wttr.in and the time tool simplified to present IANA timezone values from the agent.

## Prerequisites

- Go 1.26 or newer
- A Google API key for Gemini access
- Internet access so the app can reach:
  - Google AI endpoints
  - [wttr.in](https://wttr.in/)

## Environment setup

Set your Google API key before running the app:

```bash
export GOOGLE_API_KEY="your-google-api-key"
```

You can also place that line in your shell profile if you want it available in every session.

## Install dependencies

From the project directory:

```bash
cd /Users/jdani002/ai/adk-go/multi-tool-agent
go mod download
```

## Build the app

Build the binary:

```bash
go build ./...
```

Or build a local executable:

```bash
go build -o multi-tool-agent .
```

## Run the agent

Start the agent in interactive mode:

```bash
go run .
```

You can also use the explicit file entrypoint:

```bash
go run agent.go
```

Once it starts, enter prompts such as:

- "What is the weather in Seattle?"
- "What time is it in Europe/Paris?"
- "Compare the weather and time in Tokyo and New York"

The agent will decide which tool to use and combine the results into a helpful answer.

## Example tool behavior

- Weather requests are sent to wttr.in using the city name.
- Time requests expect a valid IANA timezone string, for example:
  - America/New_York
  - Europe/London
  - Asia/Tokyo
  - Australia/Sydney

## Launch the web UI for troubleshooting

The app includes the ADK launcher support for a web-based UI. To start it:

```bash
go run . web api webui
```

Or, equivalently:

```bash
go run agent.go web api webui
```

When the UI is running, open the local URL shown in the terminal in your browser. This is useful for:

- testing prompts interactively
- inspecting agent/tool behavior
- troubleshooting request and response issues

## Troubleshooting tips

If the app does not start correctly:

- Confirm that GOOGLE_API_KEY is set and valid.
- Make sure your machine can reach the Google AI and wttr.in endpoints.
- Run `go mod download` if dependencies are missing.
- If you are using the web UI, check the terminal output for the local address and any launcher errors.

## Notes

- Weather results are returned in the units provided by wttr.in and the agent instruction encourages the model to present them clearly.
- The current time tool uses the system's timezone database, so it can resolve common IANA zones without additional setup.
