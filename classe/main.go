// go mod init classe

package main

import (
	"classe/funcionario"
	"fmt"
)

func main() {
	individuo := funcionario.Funcionario{
		Pessoa: funcionario.Pessoa{
			Nome:  "Jordan",
			Idade: 26,
		},
		Cargo:   "devops",
		Salario: 4000.0,
	}

	fmt.Println("Nome do cara: ", individuo.Nome)

	individuo.ExibirInfo()

	individuo.Trabalhar()
}
