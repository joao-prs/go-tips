package main

import (
    "fmt"
    "log"
)

func main() {
    var num1, num2 float64
    var operation string

    // Solicita ao usuário o primeiro número
    fmt.Print("Digite o primeiro número: ")
    _, err := fmt.Scanln(&num1)
    if err != nil {
        log.Fatalf("Erro ao ler o primeiro número: %v", err)
    }

    // Solicita ao usuário a operação
    fmt.Print("Digite a operação (+, -, *, /): ")
    _, err = fmt.Scanln(&operation)
    if err != nil {
        log.Fatalf("Erro ao ler a operação: %v", err)
    }

    // Solicita ao usuário o segundo número
    fmt.Print("Digite o segundo número: ")
    _, err = fmt.Scanln(&num2)
    if err != nil {
        log.Fatalf("Erro ao ler o segundo número: %v", err)
    }

    // Realiza a operação e imprime o resultado
    switch operation {
    case "+":
        fmt.Printf("Resultado: %.2f\n", num1 + num2)
    case "-":
        fmt.Printf("Resultado: %.2f\n", num1 - num2)
    case "*":
        fmt.Printf("Resultado: %.2f\n", num1 * num2)
    case "/":
        if num2 != 0 {
            fmt.Printf("Resultado: %.2f\n", num1 / num2)
        } else {
            fmt.Println("Erro: Divisão por zero")
        }
    default:
        fmt.Println("Operação inválida")
    }
}
