package main

import (
	"fmt"
)

func main() {
	/*
		utilizando aspas, o fmt reformata corrigindo os \
	*/
	x := "Ola, bom dia.\n como vai?\tespero \"que\" tudo bem."
	fmt.Println(x)

	/*
		utilizando acento grave, ele coloca tudo literalmente sem
		corrigir os \
	*/
	y := `"Ola, bom dia.\n como vai?\tespero \"que\" tudo bem."`
	fmt.Println(y)

	z := "sad999 sadasdasd"
	fmt.Println(z)
}
