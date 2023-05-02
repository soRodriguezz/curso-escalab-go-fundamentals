/*
	Este es un ejemplo de documentación de go - usar go doc first_go_program.go
*/
package main

import "fmt"

func main() { // el compilador siempre lo busca
	// infiere que myName es de tipo string - declaración implicita
	myName := "Sebastian"
	fmt.Println("Hello, my name is", myName)
}


// ejecutar con go run first_go_program.go
// compilar con go build first_go_program.go