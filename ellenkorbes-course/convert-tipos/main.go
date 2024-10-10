package main

import (
	"fmt"
)

type hotdog int

var b hotdog = 101

func main() {
	x := 10
	fmt.Printf("%v\n", x)
	fmt.Printf("%v\n", b)

	/*
		conversao 
	*/
	x = int(b)
	fmt.Printf("%v\n", x)
}
