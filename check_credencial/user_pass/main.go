package main

import (
    "bufio"
    "fmt"
    "os"
    "path/filepath"
    "strings"
)

// Lista de arquivos ou diretórios a serem ignorados
var ignoreList = []string{
    ".terraform",
    "main.go",
}

func main() {
    directory := "./"
    foundCredentials := false
    err := filepath.Walk(directory, func(path string, info os.FileInfo, err error) error {
        if err != nil {
            return err
        }
        // Verifica se o caminho está na lista de ignorados
        for _, ignore := range ignoreList {
            if strings.Contains(path, ignore) {
                return nil
            }
        }
        // Verifica se é um arquivo regular
        if !info.IsDir() {
            if checkFileForCredentials(path) {
                foundCredentials = true
            }
        }
        return nil
    })

    if err != nil {
        fmt.Printf("[\033[91m ERROR \033[0m]\n")
        fmt.Printf("- Erro ao percorrer o diretório: %v\n", err)
    }
    if !foundCredentials {
        fmt.Printf("[\033[92m OK \033[0m]\n")
        fmt.Println("- Nenhuma credencial encontrada.")
    }
}

func checkFileForCredentials(path string) bool {
    file, err := os.Open(path)
    if err != nil {
        fmt.Printf("- Erro ao abrir o arquivo %s: %v\n", path, err)
        return false
    }
    defer file.Close()

    scanner := bufio.NewScanner(file)
    for scanner.Scan() {
        line := scanner.Text()

        if strings.Contains(line, `username: "joao"`) || strings.Contains(line, `password: "pass1"`) {
            fmt.Printf("[\033[93m FOUND \033[0m]\n")
            fmt.Printf("- Possível credencial encontrada no arquivo: %s\n", path)
            return true
        }
    }
    if err := scanner.Err(); err != nil {
        fmt.Printf("[\033[91m ERROR \033[0m]\n")
        fmt.Printf("- Erro ao ler o arquivo %s: %v\n", path, err)
    }
    return false
}
