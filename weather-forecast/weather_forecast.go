// Package weather gives weather forecasts based on location and weather condition.
package weather

// CurrentCondition takes the current weather condition which is used to forecast weather.
var CurrentCondition string

// CurrentLocation takes the current location for which the weather forcast should be generated.
var CurrentLocation string

// Forecast returns the a weather forcast based current location and current condition.
// This function takes current condition and current location and process it to generate forecast.
func Forecast(city, condition string) string {
	CurrentLocation, CurrentCondition = city, condition
	return CurrentLocation + " - current weather condition: " + CurrentCondition
}
