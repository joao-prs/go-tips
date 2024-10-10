package main

/*
	Deslocamento de bit
*/
import (
	"fmt"
)

const (
	_  = iota             // 0
	KB = 1 << (iota * 10) // 1 << (1 * 10)
	MB = 1 << (iota * 10) // 1 << (2 * 10)
	GB = 1 << (iota * 10) // 1 << (3 * 10)
	TB = 1 << (iota * 10) // 1 << (4 * 10)

	/* ou você apenas define as variaveis,
	   o iota consegue construir a sequencia
	   como mostro no exemplo abaixo.
	*/

	/*

		_  = iota             // 0
		KB = 1 << (iota * 10) // 1 << (1 * 10)
		MB
		GB
		TB

	*/
)

func main() {

	// exemplo simples, nao usa a const a cima
	a := 1
	b := a << 1
	c := b << 1
	fmt.Printf("%b\n", a)
	fmt.Printf("%b\n", b)
	fmt.Printf("%b\n", c)

	//outro exemplo, utiliza a const
	fmt.Printf("\nBinary\t\t\t\tDecimal\n")
	fmt.Printf("%b\t\t\t", KB)
	fmt.Printf("%d\n", KB)
	fmt.Printf("%b\t\t", MB)
	fmt.Printf("%d\n", MB)
	fmt.Printf("%b\t", GB)
	fmt.Printf("%d\n", GB)

}
