package main

import (
	"fmt"
	"math"
)

func main() {
	var radio float64
	const Pi = 3.14159265359

	fmt.Println("Ingrese el radio del circulo: ")
	fmt.Scanln(&radio)
	perimetro := (radio * 2) * Pi
	area := Pi*(math.Pow(radio, 2))
	fmt.Printf("El perimetro del circulo es: %.1f \n", perimetro)
	fmt.Printf("El area del circulo es: %.1f", area)

}