package helpers

import(
	"math"
	"halo/internal/app"
)

func GetRangeOfCoords(origin *app.Coordinates, radius int16) *app.RangeOfCoords {
	var ranges app.RangeOfCoords

	ranges.MaxLat = convertForLat(origin.Latitude, radius)
	ranges.MinLat = convertForLat(origin.Latitude, -1*radius)
	ranges.MaxLong = convertForLong(origin.Longitude, origin.Latitude, -1*radius)
	ranges.MinLong = convertForLong(origin.Longitude, origin.Latitude, radius)	

	return &ranges
}

func convertForLat(lat float32, radius int16) float32 {
	return lat + (float32(radius) / 69)
}

func convertForLong(long float32, lat float32, radius int16) float32 {
	return float32(long + (float32(radius) / float32(69*math.Cos(float64(lat)))))
}