package main

import (
    "fmt"
    "bufio"
    "os"
    "strconv"
    "strings"
)

// Cliente representa um cliente com nome e idade
type Cliente struct {
    Nome string
    Idade int
}



// Fila representa uma fila de clientes
type Fila struct {
    clientes []Cliente
}

// Adiciona um cliente à fila
func (f *Fila) Adicionar(cliente Cliente) {
    f.clientes = append(f.clientes, cliente)
}

// Remove e retorna o primeiro cliente da fila
func (f *Fila) Remover() Cliente {
    cliente := f.clientes[0]
    f.clientes = f.clientes[1:]
    return cliente
}

// Caixa representa um caixa que atende os clientes
type Caixa struct {
    clientesAtendidos []Cliente
}

// Atende um cliente e adiciona à lista de atendidos
func (c *Caixa) Atender(cliente Cliente) {
    c.clientesAtendidos = append(c.clientesAtendidos, cliente)
}


// LerClientesDoArquivo lê clientes de um arquivo e retorna uma fila de clientes
func LerClientesDoArquivo(nomeArquivo string) (Fila, error) {
    file, err := os.Open(nomeArquivo)
    if err != nil {
        return Fila{}, err
    }
    defer file.Close()

    fila := Fila{}
    scanner := bufio.NewScanner(file)
    for scanner.Scan() {
        linha := scanner.Text()
        partes := strings.Split(linha, ",")
        if len(partes) != 2 {
            continue // Ignora linhas malformadas
        }
        nome := strings.TrimSpace(partes[0])
        idade, err := strconv.Atoi(strings.TrimSpace(partes[1]))
        if err != nil {
            continue // Ignora idades malformadas
        }
        fila.Adicionar(Cliente{nome, idade})
    }

    if err := scanner.Err(); err != nil {
        return Fila{}, err
    }

    return fila, nil
}


func main() {
    // Lendo a fila de clientes do arquivo
    fila, err := LerClientesDoArquivo("clientes.txt")
    if err != nil {
        fmt.Println("Erro ao ler o arquivo:", err)
        return
    }

    // Criando os caixas
    caixa1 := Caixa{}
    caixa2 := Caixa{}

    // Distribuindo os clientes alternadamente entre os caixas
    caixas := []*Caixa{&caixa1, &caixa2}
    caixaAtual := 0

    for len(fila.clientes) > 0 {
        cliente := fila.Remover()
        caixas[caixaAtual].Atender(cliente)
        caixaAtual = (caixaAtual + 1) % len(caixas)
    }

    // Exibindo os clientes atendidos por cada caixa
    fmt.Println("Caixa 1 atendeu os seguintes clientes:")
    for _, cliente := range caixa1.clientesAtendidos {
        fmt.Printf("%s, %d anos\n", cliente.Nome, cliente.Idade)
    }

    fmt.Println("\nCaixa 2 atendeu os seguintes clientes:")
    for _, cliente := range caixa2.clientesAtendidos {
        fmt.Printf("%s, %d anos\n", cliente.Nome, cliente.Idade)
    }
}
