package main

import (
	"net/http"

	"github.com/linuxfreak003/weather/adapters/geo"
	httpAdapter "github.com/linuxfreak003/weather/adapters/http"
	"github.com/linuxfreak003/weather/domain"
	"github.com/linuxfreak003/weather/ports"
)

func main() {
	client := &http.Client{}
	forecaster := httpAdapter.NewWeatherGovForecaster(client)
	geocoder := geo.NewGeocoder()
	d := domain.NewDomain(forecaster, geocoder)
	d.DisplayForecast(&ports.Options{})
}
