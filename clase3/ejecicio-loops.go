package main

import (
	"fmt"
 	"math/rand" //import time package
)

func main() {
	var adivina = 99

	count := 10

	for {
		var numRandom = rand.Intn(100) + 1
		if numRandom == adivina {
			fmt.Println("Adivinaste! El numero era ", numRandom, " en ", count, " intentos")
			break
		}
		count++
		// fmt.Println(numRandom) //imprime el numero random
	}

}