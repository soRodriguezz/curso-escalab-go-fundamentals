package main

import (
	"fmt"
	"strings"
)

const (
	red = iota // pone hacia abajo 0,1,2,3...
	blue
	yellow
)

var (
	a string
	b int
	c bool
	d = "hola"
)

func main() {
	// fmt.Println("Mi IMC es", 80 / (1.78 * 1.78))
	// fmt.Print("Mi IMC es ", 80 / (1.78 * 1.78),"\n")
	// fmt.Printf("Mi %s es %v \n", "IMC",  80 / (1.78 * 1.78))

	// fmt.Println(red, blue, yellow)

	// var distancia, velocidad int
	// distancia, velocidad = 100, 50

	// var days = 28
	// var distance = 56000000
	// var totalHours = days * 24
	// fmt.Println("La velocidad necesaria para viajar a Malacandra es ", distance/totalHours, "km/h" )

	textExample := "Hola, mi nombre es Sebastian"
	word := "nombre"
	var found = strings.Contains(textExample, word) // retorna boolean si word está en textExample
	fmt.Println("¿Se encontró la palabra?", found)
}

