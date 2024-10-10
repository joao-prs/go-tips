package main

import (
	"fmt"
)
/*
	ainda não entendi bem o iota, mas é bem interessante a aplicação disso
*/
const (
	a = iota
	_ // 1
	c 
	x 
	_ // 4 
	z //= iota
)

func main() {
	
	fmt.Println(a,c,x,z)
}