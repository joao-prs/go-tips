package main

import (
	"fmt" /*int, string, float64, bool*/
)

/*
a variavel declarada em um bloco de codigo,
é undefined em outro, ou seja, se ela existe
em qualquercoisa() ela nao existe em main()
*/
var z int

func qualquercoisa(x int) {
	fmt.Println(x)
}

func main() {

	y := 10
	qualquercoisa(y)

	/* a atribuição dessa var z so pode ser feita
	dentro de um codeblock, como o caso a baixo */
	z = 10
	fmt.Println(z)

	/*
			- Decimal       %d
		    - Hexadecimal   %#x
		    - Unicode       %#U
		    - Tab           \t
		    - Linha nova    \n
	*/
	w1 := 10
	fmt.Printf("%d %#x %#U \n", w1, w1, w1)

	w2 := 18
	fmt.Printf("%d %#x %#U \n", w2, w2, w2)
}
