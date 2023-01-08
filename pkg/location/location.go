package location

import "math"

func GetDistance(lat1, lng1, lat2, lng2 float64) float64 {
	radius := 6378137.0
	pi180 := math.Pi / 180.0
	arcLat1 := lat1 * pi180
	arcLat2 := lat2 * pi180
	x := arcLat1 - arcLat2
	y := (lng1 - lng2) * pi180
	s := 2 * math.Asin(math.Sqrt(math.Pow(math.Sin(x/2), 2)+math.Cos(arcLat1)*math.Cos(arcLat2)*math.Pow(math.Sin(y/2), 2)))
	s = s * radius
	s = math.Round(s*10000) / 10000
	return s
}
