package main

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

func main() {
	//repoPath := "./" // Caminho do repositório
	repoPath := "/home/jpedro/.ssh/" // Caminho do repositório

	err := filepath.Walk(repoPath, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if !info.IsDir() {
			checkFileForKeys(path)
		}
		return nil
	})

	if err != nil {
		fmt.Printf("Erro ao percorrer o diretório: %v\n", err)
	}
}

func checkFileForKeys(filePath string) {
	file, err := os.Open(filePath)
	if err != nil {
		fmt.Printf("Erro ao abrir o arquivo %s: %v\n", filePath, err)
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		if containsKeyPattern(line) {
			fmt.Printf("Possível chave encontrada no arquivo %s: %s\n", filePath, line)
		}
	}
	if err := scanner.Err(); err != nil {
		fmt.Printf("Erro ao ler o arquivo %s: %v\n", filePath, err)
	}
}

func containsKeyPattern(line string) bool {
	keyPatterns := []string{"id_ed25519", "id_rsa", "-BEGIN OPENSSH PRIVATE KEY-", "-BEGIN RSA PRIVATE KEY-"}

	for _, pattern := range keyPatterns {
		if strings.Contains(line, pattern) {
			return true
		}
	}
	return false
}
