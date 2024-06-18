package geo

import (
	"os"

	"github.com/codingsince1985/geo-golang"
	"github.com/codingsince1985/geo-golang/arcgis"
	"github.com/codingsince1985/geo-golang/bing"
	"github.com/codingsince1985/geo-golang/chained"
	"github.com/codingsince1985/geo-golang/geocod"
	"github.com/codingsince1985/geo-golang/google"
	"github.com/codingsince1985/geo-golang/mapbox"
	"github.com/codingsince1985/geo-golang/mapquest/nominatim"
	"github.com/codingsince1985/geo-golang/mapquest/open"
	"github.com/codingsince1985/geo-golang/mapzen"
	"github.com/codingsince1985/geo-golang/opencage"
	"github.com/codingsince1985/geo-golang/openstreetmap"
	"github.com/codingsince1985/geo-golang/pickpoint"
	"github.com/codingsince1985/geo-golang/tomtom"
	"github.com/codingsince1985/geo-golang/yandex"
	"github.com/sanity-io/litter"
)

type Geocoder struct {
	geo geo.Geocoder
}

func NewGeocoder() *Geocoder {
	return &Geocoder{
		geo: chained.Geocoder(
			openstreetmap.Geocoder(),
			google.Geocoder(os.Getenv("GOOGLE_API_KEY")),
			nominatim.Geocoder(os.Getenv("MAPQUEST_NOMINATIM_KEY")),
			open.Geocoder(os.Getenv("MAPQUEST_OPEN_KEY")),
			opencage.Geocoder(os.Getenv("OPENCAGE_API_KEY")),
			bing.Geocoder(os.Getenv("BING_API_KEY")),
			google.Geocoder(os.Getenv("BAIDU_API_KEY")),
			mapbox.Geocoder(os.Getenv("MAPBOX_API_KEY")),
			openstreetmap.Geocoder(),
			pickpoint.Geocoder(os.Getenv("PICKPOINT_API_KEY")),
			arcgis.Geocoder(os.Getenv("ARCGIS_TOKEN")),
			geocod.Geocoder(os.Getenv("GEOCOD_API_KEY")),
			mapzen.Geocoder(os.Getenv("MAPZEN_API_KEY")),
			tomtom.Geocoder(os.Getenv("TOMTOM_API_KEY")),
			yandex.Geocoder(os.Getenv("YANDEX_API_KEY")),
		),
	}
}

func (g *Geocoder) GetCoordinates(address string) (float64, float64, error) {
	location, err := g.geo.Geocode(address)
	if err != nil {
		return 0, 0, err
	}
	litter.Dump(location)
	return 0, 0, nil

	//return location.Lat, location.Lng, nil
}
