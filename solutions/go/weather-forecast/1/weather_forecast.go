// Package weather is used for forecasting weather for a given city.
package weather

// CurrentCondition holds the weather condition for a specific city.
var CurrentCondition string

// CurrentLocation holds the value of the city.
var CurrentLocation string

// Forecast shows the weather condition of a specific city.
func Forecast(city, condition string) string {
	CurrentLocation, CurrentCondition = city, condition
	return CurrentLocation + " - current weather condition: " + CurrentCondition
}