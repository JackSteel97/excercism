// Package weather provides methods to get the forecast for a given city and condition.
package weather

var (
	// CurrentCondition represents the current weather condition.
	CurrentCondition string
	// CurrentLocation represents the current city location of the weather.
	CurrentLocation  string
)

// Forecast returns a string in a human/goblin-readable format representing the current weather condition at a location.
func Forecast(city, condition string) string {
	CurrentLocation, CurrentCondition = city, condition
	return CurrentLocation + " - current weather condition: " + CurrentCondition
}
