// Package weather is a tool that forecast weather condition of a current location in a city.
package weather

// CurrentCondition represents the current weather condition in a city.
var CurrentCondition string

// CurrentLocation represents the current location of a weather condition.
var CurrentLocation  string


// Forecast returns a strings of the current weather condition and location of a city.
func Forecast(city, condition string) string {
	CurrentLocation, CurrentCondition = city, condition
	return CurrentLocation + " - current weather condition: " + CurrentCondition
}
