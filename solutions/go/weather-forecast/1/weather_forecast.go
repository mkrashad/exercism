// Package weather provides tools to forecast the weather.
package weather 

var (
    // CurrentCondition represents the current weather conditon.
	CurrentCondition string
    // CurrentLocation represents the current location.
	CurrentLocation  string
)

// Forecast function takes two arguments city and weather condition and returns the string with current weather conditon and location.
func Forecast(city, condition string) string {
	CurrentLocation, CurrentCondition = city, condition
	return CurrentLocation + " - current weather condition: " + CurrentCondition
}
