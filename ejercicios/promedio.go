/*
Este programa calcula el promedio de 4 notas
se deben agregar como int o como float
*/
package main

import "fmt"

func main() {
	var nota1, nota2, nota3, nota4 float64

	fmt.Println("Digite las 4 notas")
	fmt.Scanln(&nota1)
	fmt.Scanln(&nota2)
	fmt.Scanln(&nota3)
	fmt.Scanln(&nota4)

	promedio := (nota1 + nota2 + nota3 + nota4) / 4

	fmt.Println("El promedio es:", promedio)

}