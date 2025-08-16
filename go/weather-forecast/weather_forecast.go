// Package weather forecast the current weather condition of various cities in Goblinocus.
package weather

// CurrentCondition stores the current weather condition.
var CurrentCondition string

// CurrentLocation stores the current location where we are measuring weather.
var CurrentLocation string

// Forecast() function returns the current weather forecast based on the currentLocation and CurrentCondition.
func Forecast(city, condition string) string {
	CurrentLocation, CurrentCondition = city, condition
	return CurrentLocation + " - current weather condition: " + CurrentCondition
}
