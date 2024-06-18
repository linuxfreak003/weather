package http

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"time"

	"github.com/linuxfreak003/weather/ports"
	"github.com/sanity-io/litter"
)

type weatherGovForecaster struct {
	client *http.Client
}

func NewWeatherGovForecaster(client *http.Client) *weatherGovForecaster {
	return &weatherGovForecaster{
		client: client,
	}
}

func getURL(url string) ([]byte, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, fmt.Errorf("error getting url: %s", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		return nil, fmt.Errorf("error getting url: %s", resp.Status)
	}

	buffer := bytes.NewBuffer(nil)

	_, err = io.Copy(buffer, resp.Body)
	if err != nil {
		return nil, fmt.Errorf("error copying response body: %s", err)
	}

	return buffer.Bytes(), nil
}

func getApiForecast(url string) ([]byte, error) {
	bytes, err := getURL(url)
	if err != nil {
		return nil, err
	}

	r := map[string]interface{}{}
	json.Unmarshal(bytes, &r)

	url = r["properties"].(map[string]interface{})["forecast"].(string)

	bytes, err = getURL(url)
	if err != nil {
		return nil, err
	}

	return bytes, nil
}

func (w *weatherGovForecaster) GetForecast(longitude, latitude float64) (*ports.Forecast, error) {
	url := fmt.Sprintf("https://api.weather.gov/points/%0.4f,%0.4f", longitude, latitude)
	bs, err := getApiForecast(url)
	if err != nil {
		return nil, fmt.Errorf("could not get forcast: %s", err)
	}

	fmt.Println(string(bs))
	var res Response
	err = json.Unmarshal(bs, &res)
	if err != nil {
		log.Panicf("Could not unmarshal json: %v", err)
	}

	fmt.Println(res)
	litter.Dump(res)

	return &ports.Forecast{
		Elevation:     res.Properties.Elevation.Value,
		ElevationUnit: res.Properties.Elevation.UnitCode,
		Periods: func([]Period) (periods []*ports.Period) {
			for _, p := range res.Properties.Periods {
				periods = append(periods, &ports.Period{
					Name:            p.Name,
					Temperature:     p.Temperature,
					TemperatureUnit: p.TemperatureUnit,
					Sky:             p.ShortForecast,
					Wind: &ports.Wind{
						Speed:     p.WindSpeed,
						Direction: p.WindDirection,
					},
				})
			}
			return
		}(res.Properties.Periods),
	}, nil
}

type Response struct {
	Properties Properties `json:"properties"`
}

type Elevation struct {
	Value    float64 `json:"value"`
	UnitCode string  `json:"unitCode"`
}

type Period struct {
	Number           int       `json:"number"`
	Name             string    `json:"name"`
	StartTime        time.Time `json:"startTime"`
	EndTime          time.Time `json:"endTime"`
	IsDaytime        bool      `json:"isDaytime"`
	Temperature      int       `json:"temperature"`
	TemperatureUnit  string    `json:"temperatureUnit"`
	WindSpeed        string    `json:"windSpeed"`
	WindDirection    string    `json:"windDirection"`
	Icon             string    `json:"icon"`
	ShortForecast    string    `json:"shortForecast"`
	DetailedForecast string    `json:"detailedForecast"`
}

type Properties struct {
	Updated           string    `json:"updated"`
	Units             string    `json:"units"`
	ForecastGenerator string    `json:"forecastGenerator"`
	GeneratedAt       time.Time `json:"generatedAt"`
	UpdateTime        time.Time `json:"updateTime"`
	Elevation         Elevation `json:"elevation"`
	Periods           []Period  `json:"periods"`
}
