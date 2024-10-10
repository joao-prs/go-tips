package main

import (
	"fmt"
)

// o tipo constante nao é definido
const x = 10
// o tipo variavel é definido 
var y = 10

var f float64

func main() {
	
	f = y
	fmt.Printf(f)
}