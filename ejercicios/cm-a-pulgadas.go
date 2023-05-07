package main

import "fmt"

func main() {
	var centimetros float64
	const CmPorPulgada = 2.54

	fmt.Println("Ingrese la cantidad de centimetros que desea convertir a pulgadas:")
	fmt.Scanf("%f", &centimetros)
	pulgadas := centimetros / CmPorPulgada
	fmt.Printf("%.f cm equivalen a %.4f pulgadas\n", centimetros, pulgadas)

}