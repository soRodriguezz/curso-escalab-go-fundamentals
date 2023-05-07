package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {

	// ! Este codigo solo imprime Sebastian
	// var input string
	// fmt.Println("Ingrese su nombre: ")
	// fmt.Scanln(&input)
	// fmt.Println("Hola, ", input)

	fmt.Println("Ingrese su nombre: ")
	scanner := bufio.NewScanner(os.Stdin)
	if scanner.Scan() {
		input := scanner.Text()
		fmt.Println("Hola,", input)
	}
}