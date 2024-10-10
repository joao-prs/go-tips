package main

import (
	"fmt"
)

// Cap. 6 – Fluxo de Controle – 5
/*
	o break literalmente quebra uma estrutura
	de repetição de qualquer ponto dentro dele
*/
func main() {
	fmt.Printf(" - Fluxo de controle\n")

	// break example
	fmt.Println("\nx: ")

	for x := 0; x < 10; x++ {
		fmt.Print(" ", x)
		if x == 6 {
			break
		}
	}

	// end code
	print("\n")
}
