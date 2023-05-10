package main

import "fmt"

func celsiusToFahrenheit(c float64) float64 {
	return (c * 9.0 / 5.0) + 32.0
}

func kelvinToCelsius(k float64) float64 {
	celsius := celsiusToFahrenheit(kelvinToCelsius(k))
	return fahrenheit
}

func main() {
	//kelvin := 233.0
	//celsius := kelvinToCelsius(kelvin)
	//fmt.Print(kelvin, "o K is ", celsius, "o C")

	celsius := 10.0
	fahrenheit := celsiusToFahrenheit(celsius)
	fmt.Print(celsius, "ºC is ", fahrenheit, "ºF")
}


