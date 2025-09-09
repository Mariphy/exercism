//Package weather provides tools to forecast weather.
package weather

//CurrentCondition represents current weather conditions.
var CurrentCondition string
//CurrentLocation represents current location.
var CurrentLocation string

//Forecast accepts city and weather condition and returns forecast for current location.
func Forecast(city, condition string) string {
	CurrentLocation, CurrentCondition = city, condition
	return CurrentLocation + " - current weather condition: " + CurrentCondition
}
