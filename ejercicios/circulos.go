package main

import (
	"fmt"
	"math"
)

func main() {
	var radio float64
	const PI = 3.14159265359

	fmt.Println("Ingrese el radio del circulo: ")
	fmt.Scanln(&radio)
	perimetro := (radio * 2) * PI
	area := PI*(math.Pow(radio, 2))
	fmt.Printf("El perimetro del circulo es: %.1f \n", perimetro)
	fmt.Printf("El area del circulo es: %.1f", area)

}