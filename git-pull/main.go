package main

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

// Cores
const (
	VML = "\033[0;31m"
	AML = "\033[93m"
	ROX = "\033[95m"
	VRD = "\033[0;32m"
	RST = "\033[0m"
)

func runCommand(name string, args ...string) error {
	cmd := exec.Command(name, args...)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	return cmd.Run()
}

func main() {
	// Diretório onde estão os repositórios
	repositoriosDir, err := os.Getwd()
	if err != nil {
		fmt.Println(string(VML), "Erro ao obter o diretório atual.", string(RST))
		return
	}

	// Verifica se o diretório existe
	if _, err := os.Stat(repositoriosDir); os.IsNotExist(err) {
		fmt.Println(string(VML), "O diretório dos repositórios não foi encontrado.", string(RST))
		return
	}

	// Navega até o diretório dos repositórios
	err = os.Chdir(repositoriosDir)
	if err != nil {
		fmt.Println(string(VML), "Erro ao navegar até o diretório dos repositórios.", string(RST))
		return
	}

	// Loop sobre todas as pastas no diretório
	dirs, err := os.ReadDir(repositoriosDir)
	if err != nil {
		fmt.Println(string(VML), "Erro ao ler diretórios.", string(RST))
		return
	}

	for _, dir := range dirs {
		if dir.IsDir() {
			repoPath := filepath.Join(repositoriosDir, dir.Name())
			if _, err := os.Stat(filepath.Join(repoPath, ".git")); err == nil {
				// Executar dentro de cada repo
				err := os.Chdir(repoPath)
				if err != nil {
					fmt.Println(string(VML), "Erro ao entrar no repositório:", repoPath, string(RST))
					continue
				}

				fmt.Printf("%sAtualizando repositório:%s %s%s\n", ROX, AML, dir.Name(), RST)

				// Loop sobre todas as branches
				branchesOutput, err := exec.Command("git", "branch", "-l").Output()
				if err != nil {
					fmt.Println(string(VML), "Erro ao listar branches no repositório:", repoPath, string(RST))
					continue
				}

				branches := strings.Split(string(branchesOutput), "\n")
				for _, branch := range branches {
					branch = strings.TrimSpace(branch)
					branch = strings.TrimPrefix(branch, "* ")
					if branch == "" {
						continue
					}
					//fmt.Printf("Fazendo checkout da branch: %s\n", branch)
					err = runCommand("git", "checkout", branch)
					if err != nil {
						fmt.Printf("%sErro ao fazer checkout da branch: %s%s\n", VML, branch, RST)
						continue
					}
					//fmt.Printf("Fazendo pull na branch: %s\n", branch)
					err = runCommand("git", "pull")
					if err != nil {
						fmt.Printf("%sErro ao fazer pull na branch: %s%s\n", VML, branch, RST)
					}
				}

				// Retorne ao diretório principal
				err = os.Chdir(repositoriosDir)
				if err != nil {
					fmt.Println(string(VML), "Erro ao retornar ao diretório principal.", string(RST))
					return
				}
			}
		}
	}

	// Retorne com êxito
	fmt.Println(string(VRD), "Atualização concluída.", string(RST))
}
