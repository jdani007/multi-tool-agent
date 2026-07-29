package data

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"google.golang.org/adk/v2/agent"
)

type weatherReport struct {
	CurrentCondition []struct {
		FeelsLike      string `json:"FeelsLikeF"`
		Cloudcover      string `json:"cloudcover"`
		Humidity        string `json:"humidity"`
		ObservationTime string `json:"observation_time"`
		Precipitation    string `json:"precipInches"`
		Pressure  string `json:"pressureInches"`
		Temperature           string `json:"temp_F"`
		UvIndex         string `json:"uvIndex"`
		Visibility string `json:"visibilityMiles"`
		WeatherCode     string `json:"weatherCode"`
		WeatherDescription     []struct {
			Value string `json:"value"`
		} `json:"weatherDesc"`
		WeatherIconURL []struct {
			Value string `json:"value"`
		} `json:"weatherIconUrl"`
		Winddir16Point string `json:"winddir16Point"`
		WinddirDegree  string `json:"winddirDegree"`
		Windspeed string `json:"windspeedMiles"`
	} `json:"current_condition"`
	NearestArea []struct {
		City []struct {
			Value string `json:"value"`
		} `json:"areaName"`
		Country []struct {
			Value string `json:"value"`
		} `json:"country"`
		Latitude   string `json:"latitude"`
		Longitude  string `json:"longitude"`
		Population string `json:"population"`
		Region     []struct {
			Value string `json:"value"`
		} `json:"region"`
		WeatherURL []struct {
			Value string `json:"value"`
		} `json:"weatherUrl"`
	} `json:"nearest_area"`
	Request []struct {
		Query string `json:"query"`
		Type  string `json:"type"`
	} `json:"request"`
	Weather []struct {
		Astronomy []struct {
			MoonIllumination string `json:"moon_illumination"`
			MoonPhase        string `json:"moon_phase"`
			Moonrise         string `json:"moonrise"`
			Moonset          string `json:"moonset"`
			Sunrise          string `json:"sunrise"`
			Sunset           string `json:"sunset"`
		} `json:"astronomy"`
		Averagetemperature    string `json:"avgtempF"`
		Date        string `json:"date"`
		MaximumTemperature    string `json:"maxtempF"`
		MinimumTemperature    string `json:"mintempF"`
		SunHour     string `json:"sunHour"`
		TotalSnow string `json:"totalSnow_cm"`
		UvIndex     string `json:"uvIndex"`
	} `json:"weather"`
	Status       string
	ErrorMessage string
}


func GetWeather(_ agent.Context, args inputArgs) (weatherReport, error) {
	report, err := getExternalWeather(args.City)
	if err != nil {
		return weatherReport{}, fmt.Errorf("unable to get weather report for %v: %v", args.City, err)
	}
	return report, nil
}


func getExternalWeather(city string) (weatherReport, error) {
	url := fmt.Sprintf("https://wttr.in/%v?format=j2", city)
	resp, err := http.Get(url)
	if err != nil {
		return weatherReport{}, err
	}
	defer resp.Body.Close()

	body,  err := io.ReadAll(resp.Body)
	if err != nil {
		return weatherReport{}, nil
	}

	var report weatherReport
	if err := json.Unmarshal(body, &report); err != nil {
		return weatherReport{}, err
	}

	return report, nil
}