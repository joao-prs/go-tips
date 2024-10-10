package main

import (
	"fmt"
	"net"
	"time"
)

func umaortaUnica(host string, porta int) {
	endereco := fmt.Sprintf("%s:%d", host, porta)
	conn, err := net.DialTimeout("tcp", endereco, 1*time.Second)
	if err != nil {
		fmt.Printf("port %d CLOSED\n", porta)
		return
	}
	conn.Close()
	fmt.Printf("port %d OK\n", porta)
	return
}



func rPorta(host string, porta int) bool {
	endereco := fmt.Sprintf("%s:%d", host, porta)
	conn, err := net.DialTimeout("tcp", endereco, 1*time.Second)
	if err != nil {
		return false
	}
	conn.Close()
	return true
}

func rangePortas(host string, portaInicial, portaFinal int) {
	contadorPortas := 0
	resultados := make(map[int]bool)
	for porta := portaInicial; porta <= portaFinal; porta++ {
		// guarda em um array
		resultados[porta] = rPorta(host, porta)
	}

	// Contar as portas abertas
	for _, aberto := range resultados {
		if aberto {
			contadorPortas++
		}
	}

	fmt.Printf("Existem %d portas abertas\n", contadorPortas)
	fmt.Printf("A porta abertas sao: \n")

	// Listar as portas abertas
	for porta, aberto := range resultados {
		if aberto {
			fmt.Printf("%d ", porta)
		}
	}
	fmt.Println()
}
