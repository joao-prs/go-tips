// example print bruteforce

package main

import (
	"fmt"
)

const (
	charset = "0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZ!@#$%^&*()-_=+[]{}|;:'\",.<>/?`~" // Conjunto de caracteres
)

func generateCombinations(prefix string, length int) {
	if length == 0 {
		fmt.Println(prefix)
		return
	}
	
	// Gera combinações chamando a função recursivamente para cada caractere
	for i := 0; i < len(charset); i++ {
		generateCombinations(prefix+string(charset[i]), length-1)
	}
}

func main() {
	length := 1
	for {
		generateCombinations("", length)
		length++ // Aumenta o comprimento para gerar combinações maiores
	}
}