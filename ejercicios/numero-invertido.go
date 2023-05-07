package main

import (
	"fmt"
)

func main() {
	var numero string

	fmt.Println("Ingrese un numero")
	fmt.Scanln(&numero)
	runas := []rune(numero)

	for i := 0; i < len(runas)/2; i++ {
		j := len(runas) - i - 1
		runas[i], runas[j] = runas[j], runas[i]
	}
	invertida := string(runas)

	fmt.Println(invertida)
}
