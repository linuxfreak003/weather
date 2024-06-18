package domain

import (
	"fmt"

	"github.com/linuxfreak003/weather/ports"
)

type Domain struct {
	forecaster ports.Forecaster
	geocoder   ports.Geocoder
}

func NewDomain(f ports.Forecaster, g ports.Geocoder) ports.Domain {
	return Domain{
		forecaster: f,
		geocoder:   g,
	}
}

func (d Domain) DisplayForecast(opts *ports.Options) {
	address := fmt.Sprintf("%s, %s, %s, %s", opts.City, opts.State, opts.Country, opts.Postal)
	lat, lng, err := d.geocoder.GetCoordinates(address)
	if err != nil {
		fmt.Println("Could not get coordinates!")
		return
	}

	forecast, err := d.forecaster.GetForecast(lng, lat)
	if err != nil {
		fmt.Println("Could not get forcast!")
		return
	}

	fmt.Println(forecast)
}
