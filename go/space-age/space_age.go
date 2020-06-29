package space

// Planet type for name
type Planet string

const earthYear = 31557600.0

// Age returns the age in Planet years
func Age(seconds float64, planet Planet) (age float64) {
	relativeEarthYear := map[Planet]float64{
		"Mercury": 0.2408467,
		"Venus":   0.61519726,
		"Earth":   1.0,
		"Mars":    1.8808158,
		"Jupiter": 11.862615,
		"Saturn":  29.447498,
		"Uranus":  84.016846,
		"Neptune": 164.79132,
	}

	return seconds / (earthYear * relativeEarthYear[planet])
}
