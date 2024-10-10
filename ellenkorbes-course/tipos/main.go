package main

import (
	"fmt"
)

/*
	como criar o proprio tipo
*/

type hotdog int

var b hotdog
var c hotdog = 10

func main() {
	x := 10
	fmt.Printf("%T\n", x)
	fmt.Printf("%T\n", b)

	/*
		aparentemente sao os mesmos valores, mas nao
		sao o mesmo tipo.
	*/
	fmt.Printf("%v\n", x)
	fmt.Printf("%v\n", c)
}
