
package main

import "fmt"

func main() {
	host := "localhost"
	fmt.Printf("host: %s\n\n", host)

	porta:= 22
	umaortaUnica(host, porta)
	
	fmt.Printf("Kubernetes API server:\n")
	porta= 80
	umaortaUnica(host, porta)

	// Testar um intervalo de portas
	fmt.Printf("NodePort Services:\n")
	portaInicial := 10
	portaFinal := 60
	rangePortas(host, portaInicial, portaFinal)
}

// go run main.go funcs.go