
package main

import (
	"fmt"
 	"time" //import time package
)

func main() {
	var count = 10
	for count > 0 {
		fmt.Println(count)
		time.Sleep(time.Second) //sleep for 1 second
		count--
	}
	fmt.Println("Buum!")
 }
 