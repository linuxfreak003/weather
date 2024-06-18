package ports

type Domain interface {
	DisplayForecast(*Options)
}

type Options struct {
	City      string
	State     string
	Country   string
	Postal    string
	Latitude  float64
	Longitude float64
}
