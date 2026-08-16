// Package weather provides tools to fetch the current weather condition for a given city.
package weather

var (
	// CurrentCondition is a string representing the current weather condition.
	CurrentCondition string
	// CurrentLocation is a string representing the name of a city.
	CurrentLocation string
)

// Forecast returns a string describing the weather condition in a given city. It takes the name of a city and its current weather condition as arguments and outputs a simple weather forecast.
func Forecast(city, condition string) string {
	CurrentLocation, CurrentCondition = city, condition
	return CurrentLocation + " - current weather condition: " + CurrentCondition
}
