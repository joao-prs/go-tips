package main

import (
	"fmt"
)

// Cap. 6 – Fluxo de Controle – 2 a 4
/*
	na linguagem go nao existe while, entao
	sao "for" que resolvem esse problema de
	formas... diferentes.
*/

func main() {
	fmt.Printf(" - Fluxo de controle\n")

	// simple for
	fmt.Println("\nx: ")

	for x := 0; x < 10; x++ {
		fmt.Print(" ", x)
	}

	// while
	fmt.Println("\ny: ")

	y := 0

	for y < 10 {
		fmt.Print(" ", y)
		y++
	}

	// infinite
	fmt.Println("\ny: ")

	z := 0

	for {
		if z < 10 {
			fmt.Print(" ", z)
			z++
		} else {
			// o que quebra o loop
			fmt.Print(" cabou")
			break
		}
	}

	// end code
	print("\n")
}
