package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"os"
	"strings"

	"github.com/theritikchoure/logx"
)

const (
	API_URL  = "https://api.weatherapi.com/v1/current.json"
	ENV_NAME = "DEFAULT_WEATHER_LOCATION"
	API_KEY  = "587872da62b842af94a152955213010"
)

type Location struct {
	City    string
	Region  string
	Country string
}

type Weather struct {
	Location            Location
	Temperature         float64
	FeelsLike           float64
	Humidity            float64
	WindSpeed           float64
	PrecipitationChance float64
}

func (w Weather) String() string {
	return fmt.Sprintf(
		"Temperature: %.2f°C, Location: %s, Feels Like: %.2f°C, Humidity: %.2f%%, Wind Speed: %.2f km/h, Precipitation Chance: %.2f%%",
		w.Temperature,
		w.Location.City,
		w.FeelsLike,
		w.Humidity,
		w.WindSpeed,
		w.PrecipitationChance,
	)
}

//	func (weather Weather) Print() {
//		// Determine the color for the temperature based on its value
//		var temperatureColor string
//		switch {
//		case weather.Temperature < 20.0:
//			temperatureColor = logx.FGBLUE
//		case weather.Temperature < 30.0:
//			temperatureColor = logx.FGGREEN
//		case weather.Temperature < 40.0:
//			temperatureColor = logx.FGYELLOW
//		default:
//			temperatureColor = logx.FGRED
//		}
//
//		// Print temperature with color, then other weather details
//		logx.Logf("Temperature: %.2f°C", temperatureColor, "", weather.Temperature)
//		fmt.Printf("%s, %s, %s; %.2f°C; %.2f%%; %.2f km/h; %.2f%%\n",
//			weather.Location.City, weather.Location.Region, weather.Location.Country, weather.FeelsLike, weather.Humidity, weather.WindSpeed, weather.PrecipitationChance)
//	}

func (weather Weather) Print() {
	var temperatureColor string
	var emoji string
	if weather.Temperature < 20.0 {
		temperatureColor = logx.FGBLUE
		emoji = "🥶"
	} else if weather.Temperature < 30.0 {
		temperatureColor = logx.FGGREEN
		emoji = "😎"
	} else if weather.Temperature < 40.0 {
		temperatureColor = logx.FGYELLOW
		emoji = "🥺"
	} else {
		temperatureColor = logx.FGRED
		emoji = "🥵"
	}

	logx.Logf("%s %.2f°C, feeks like %.2f°C", temperatureColor, "", emoji, weather.Temperature, weather.FeelsLike)
	fmt.Printf("📍 %s, %s, %s\n", weather.Location.City, weather.Location.Region, weather.Location.Country)
	// logx.Logf("Feels Like: %.2f°C %s", "", "", weather.FeelsLike, emoji)
	//
	// fmt.Printf("Humidity: %.2f%% 💧\n", weather.Humidity)
	// fmt.Printf("Wind Speed: %.2f km/h 💨\n", weather.WindSpeed)
	// fmt.Printf("Precipitation Chance: %.2f%% 🌧️\n", weather.PrecipitationChance)
}

func GetWeather(location string) (*Weather, error) {
	apiUrl := fmt.Sprintf("%s?key=%s&q=%s&aqi=no", API_URL, API_KEY, url.QueryEscape(location))

	// Make a GET request to the API URL
	resp, err := http.Get(apiUrl)
	if err != nil {
		printError(fmt.Errorf("failed to make GET request to %s: %w", apiUrl, err))
		return nil, err
	}

	var weatherResponse WeatherResponse
	err = json.NewDecoder(resp.Body).Decode(&weatherResponse)
	if err != nil {
		printError(fmt.Errorf("failed to decode JSON response: %w", err))
		return nil, err
	}

	if weatherResponse.Error.Code != 0 {
		printError(fmt.Errorf("%s, searching for %s", weatherResponse.Error.Message, location))
		return nil, fmt.Errorf("API error: %s", weatherResponse.Error.Message)
	}

	weather := Weather{
		Location: Location{
			City:    weatherResponse.Location.Name,
			Region:  weatherResponse.Location.Region,
			Country: weatherResponse.Location.Country,
		},
		Temperature:         weatherResponse.Current.TempC,
		FeelsLike:           weatherResponse.Current.FeelslikeC,
		Humidity:            float64(weatherResponse.Current.Humidity),
		WindSpeed:           weatherResponse.Current.WindMph,
		PrecipitationChance: weatherResponse.Current.PrecipMm,
	}

	return &weather, nil
}

func printHelp() {
	fmt.Println("Usage: weather [location]")
	fmt.Printf("If no location is provided, the default location will be used. Set the %s environment variable to change the default location.\n", ENV_NAME)
}

func printError(err error) {
	logx.Logf("Error: %s", logx.FGRED, "", err)
}

func main() {
	var location string

	if len(os.Args) > 1 {
		location = strings.Join(os.Args[1:], " ")
	} else {
		defaultLocaton := os.Getenv(ENV_NAME)

		if defaultLocaton == "" {
			logx.Logf("No default location found. Please set the environment variable %s", "", logx.FGYELLOW, ENV_NAME)
			printHelp()
			return
		}
		location = defaultLocaton
	}

	weather, err := GetWeather(location)
	if err != nil {
		return
	}

	weather.Print()
}

type WeatherResponse struct {
	Location LocationResponse `json:"location"`
	Current  Current          `json:"current"`
	Error    APIError         `json:"error"`
}

type APIError struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

type LocationResponse struct {
	Name           string  `json:"name"`
	Region         string  `json:"region"`
	Country        string  `json:"country"`
	Lat            float64 `json:"lat"`
	Lon            float64 `json:"lon"`
	TzID           string  `json:"tz_id"`
	LocaltimeEpoch int64   `json:"localtime_epoch"`
	Localtime      string  `json:"localtime"`
}

type Current struct {
	LastUpdatedEpoch int64     `json:"last_updated_epoch"`
	LastUpdated      string    `json:"last_updated"`
	TempC            float64   `json:"temp_c"`
	TempF            float64   `json:"temp_f"`
	IsDay            int       `json:"is_day"`
	Condition        Condition `json:"condition"`
	WindMph          float64   `json:"wind_mph"`
	WindKph          float64   `json:"wind_kph"`
	WindDegree       int       `json:"wind_degree"`
	WindDir          string    `json:"wind_dir"`
	PressureMb       float64   `json:"pressure_mb"`
	PressureIn       float64   `json:"pressure_in"`
	PrecipMm         float64   `json:"precip_mm"`
	PrecipIn         float64   `json:"precip_in"`
	Humidity         int       `json:"humidity"`
	Cloud            int       `json:"cloud"`
	FeelslikeC       float64   `json:"feelslike_c"`
	FeelslikeF       float64   `json:"feelslike_f"`
	WindchillC       float64   `json:"windchill_c"`
	WindchillF       float64   `json:"windchill_f"`
	HeatindexC       float64   `json:"heatindex_c"`
	HeatindexF       float64   `json:"heatindex_f"`
	DewpointC        float64   `json:"dewpoint_c"`
	DewpointF        float64   `json:"dewpoint_f"`
	VisKm            float64   `json:"vis_km"`
	VisMiles         float64   `json:"vis_miles"`
	Uv               float64   `json:"uv"`
	GustMph          float64   `json:"gust_mph"`
	GustKph          float64   `json:"gust_kph"`
}

type Condition struct {
	Text string `json:"text"`
	Icon string `json:"icon"`
	Code int    `json:"code"`
}
