package data

import (
	"fmt"
	"strings"
	"time"

	"google.golang.org/adk/v2/agent"
)

func GetCurrentTime(_ agent.Context, args inputArgs) (response, error) {
	city := strings.ToLower(args.City)

	switch city {
	case "nyc", "new york", "new york city":
		tz, err := time.LoadLocation("America/New_York")
		if err != nil {
			return response{}, err
		}
		now := time.Now().In(tz)

		return response{
			Status: "success",
			Report: fmt.Sprintf("The current time in %v is %v", args.City, now.Format("2006-01-02 15:04:05 MST-0700")),
		}, nil
	default:
		return response{
			Status:       "error",
			ErrorMessage: fmt.Sprintf("Sorry, I don't have timezone information for %v.", args.City),
		}, nil

	}
}
