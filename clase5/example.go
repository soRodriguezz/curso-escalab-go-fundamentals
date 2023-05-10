package main

import "fmt"

type celsius float64
type kelvin float64
type  fahrenheit float64

func (c celsius) fahrenheit() fahrenheit {
	return fahrenheit(c * 9.0 / 5.0 + 32.0)
}

func main() {
	c.fahrenheit()
}