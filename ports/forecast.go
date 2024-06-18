package ports

import (
	"fmt"
	"strings"
)

var ArrowMap = map[string]string{
	"NW": "⬁",
	"NE": "⬀",
	"SE": "⬂",
	"SW": "⬃",
	"N":  "⇧",
	"S":  "⇩",
	"E":  "⇨",
	"W":  "⇦",
}

type Forecast struct {
	Address       string
	Elevation     float64
	ElevationUnit string
	Periods       []*Period
}

type Period struct {
	Name            string
	Temperature     int
	TemperatureUnit string
	Sky             string
	Wind            *Wind
}

type Wind struct {
	Speed     string
	Direction string
}

type StringBuilder struct {
	s strings.Builder
}

func NewStringBuilder() *StringBuilder {
	return &StringBuilder{
		s: strings.Builder{},
	}
}

func (b *StringBuilder) Writef(format string, args ...interface{}) {
	b.s.WriteString(fmt.Sprintf(format, args...))
}

func (b *StringBuilder) String() string {
	return b.s.String()
}

func (f *Forecast) String() string {
	b := NewStringBuilder()

	b.Writef("Elevation: %0.2f %s\n",
		f.Elevation,
		f.ElevationUnit,
	)
	for _, period := range f.Periods {
		b.Writef("%s:\n", period.Name)
		b.Writef("\tTemperature: %d %s\n", period.Temperature, period.TemperatureUnit)
		b.Writef("\tSky: %s\n", period.Sky)
		b.Writef("\tWind: %s to the %s %s\n", period.Wind.Speed, period.Wind.Direction, ArrowMap[period.Wind.Direction])
	}
	return b.String()
}
