package main

import (
	"fmt"
 	"math/rand" //import time package
)

func main() {
	var advina = 14

	count := 10

	for count > 0 {
		var numRandom = rand.Intn(100) + 1
		fmt.Println(numRandom)
		if numRandom == advina {
			break
		}
		count--
	}

}