package main

import (
    "fmt"
)

func main() {
    var name string

    // Solicita ao usuário que digite o seu nome
    fmt.Print("Digite seu nome: ")
    fmt.Scanln(&name)

    // Exibe a mensagem de boas-vindas
    fmt.Printf("Bem-vindo %s!\n", name)
}
