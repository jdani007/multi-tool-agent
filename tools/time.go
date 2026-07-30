package tools

import (
	"time"

	"google.golang.org/adk/v2/agent"
)

type timeInput struct {
	IANATimeZone string
}

type timeOutput struct {
	Time       string
}

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
