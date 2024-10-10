package funcionario

import (
	"fmt"
)

// Pessoa
type Pessoa struct {
	Nome  string
	Idade int
}

// Método de pessoa
func (p *Pessoa) ExibirInfo() {
	fmt.Println("Nome:", p.Nome)
	fmt.Println("Idade:", p.Idade)
}

// funcionario
type Funcionario struct {
	Pessoa
	Cargo   string
	Salario float32
}

// Método de funcionario
func (f *Funcionario) Trabalhar() {
	fmt.Printf("\ntrabalhando...\n")
}
