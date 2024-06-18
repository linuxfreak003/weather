package ports

type Forecaster interface {
	GetForecast(longitude, latitude float64) (*Forecast, error)
}
