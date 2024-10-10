package main

import (
	"fmt"
)
/*
	formas de declarar variaveis dentro e fora
	do escopo, usando a marmota (:=)
*/

// forma certa de declarar
var y = "hola"

// forma errada
// y := "hola"


func main() {
	fmt.Println(y)

	z1 := "2"
	z2 := 15
	fmt.Println(z1, z2)

	// statement (declaração, afirmação)
	x := 10 < 20
	fmt.Println(x)
	// true
}