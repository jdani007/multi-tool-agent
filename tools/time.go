package tools

import (
	"time"

	"google.golang.org/adk/v2/agent"
)

// timeInput holds the IANA time zone identifier for the requested location.
type timeInput struct {
	IANATimeZone string
}

// timeOutput contains the formatted time returned to the agent.
type timeOutput struct {
	Time string
}

// GetCurrentTime returns the current local time for the requested time zone.
func GetCurrentTime(_ agent.Context, args timeInput) (timeOutput, error) {

	timeZone, err := time.LoadLocation(args.IANATimeZone)
	if err != nil {
		return timeOutput{}, err
	}
	t := time.Now().In(timeZone).Format("2006-01-02 15:04:05 MST-0700")

	return timeOutput{
		Time: t,
	}, nil
}
