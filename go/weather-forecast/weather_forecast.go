// Package weather provides function related to forcasting the
// weather.
package weather

// CurrentCondition is a string representing the current weather condition.
var CurrentCondition string
// CurrentLocation is a string representing the current weather location.
var CurrentLocation string

// Forecast takes the city and current weather conditions and returns a
// formatted string explaining the weather.
func Forecast(city, condition string) string {
	CurrentLocation, CurrentCondition = city, condition
	return CurrentLocation + " - current weather condition: " + CurrentCondition
}
