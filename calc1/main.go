package main

import (
	"fmt"
)

func main() {

	// questao de uma prova(mas nao era em golang)
	a := 5
	b := 10
	c := 3
	var x int

	x = a * (b + c)
	x++

	if x > c {
		x = x + (b + (a - c))
	} else {
		x = a - (c * b)
	}

	x = x - (x / 2)

	if a < b {
		x = x - (a + b)
	} else {
		x = x - (a - b)
	}

	x++

	fmt.Printf("%v\n", x)
}
