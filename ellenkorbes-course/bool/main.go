package main

import (
	"fmt"
)

var x bool

func main() {
	// exemplo de bool simples
	fmt.Println(x)
	x = true
	fmt.Println(x)

	// parenteses sao opcionais
	y := (10 < 100)
	z := 10 > 100

	fmt.Println(y)
	fmt.Println(z)
}