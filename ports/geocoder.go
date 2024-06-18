package ports

type Geocoder interface {
	//GetAddress(longitude, latitude float64) (string, error)
	GetCoordinates(address string) (float64, float64, error)
}
