package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"os/exec"
	"runtime"
	"strconv"
	"strings"
	"time"
)

func clear_screen() {
	var cmd *exec.Cmd
	switch runtime.GOOS {
	case "windows":
		cmd = exec.Command("cmd", "/c", "cls")
	default:
		cmd = exec.Command("clear")
	}
	cmd.Stdout = os.Stdout
	cmd.Run()
}

func confirma_enter() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("\r\033[0;33mprosseguir? [enter]: \033[0m")
	confirm, _ := reader.ReadString('\n')
	confirm = strings.TrimSpace(confirm)

	if confirm == "" {
		confirm = "y"
	}
}

func main() {
	var array []string
	reader := bufio.NewReader(os.Stdin)

	clear_screen()
	for {
		fmt.Println(` _____       github@joao-prs`)
		fmt.Println(`|A .  | _____               `)
		fmt.Println(`| /.\ ||A ^  | _____        `)
		fmt.Println(`|(_._)|| / \ ||A _  | _____ `)
		fmt.Println(`|  |  || \ / || ( ) ||A_ _ |`)
		fmt.Println(`|____V||  .  ||(_'_)||( v )|`)
		fmt.Println(`       |____V||  |  || \ / |`)
		fmt.Println(`              |____V||  .  |`)
		fmt.Println(`                     |____V|`)

		fmt.Println("\nEscolha uma opção:")
		fmt.Println("1. Adicionar cartas ao baralho")
		fmt.Println("2. Listar deck de cartas")
		fmt.Println("3. Embaralhar o deck")
		fmt.Println("4. Sair do programa\n")

		input, _ := reader.ReadString('\n')
		choice := strings.TrimSpace(input)

		switch choice {
		case "1":
			fmt.Println("\nQuantas cartas você deseja adicionar?")
			input, _ = reader.ReadString('\n')
			count, err := strconv.Atoi(strings.TrimSpace(input))
			if err != nil {
				clear_screen()
				fmt.Println("\033[1;91mEntrada inválida, por favor insira um número!\033[0m")
				continue
			}

			for i := 0; i < count; i++ {
				fmt.Printf("Digite a carta %d: ", i+1)
				element, _ := reader.ReadString('\n')
				array = append(array, strings.TrimSpace(element))
			}
			fmt.Println("\nFim da lista!")

			// confirmacao
			confirma_enter()
			clear_screen()

		case "2":
			fmt.Println("\nlista de cartas:")
			for i, element := range array {
				fmt.Printf("%d:\033[1;96m %s\n\033[39m", i+1, element)
			}
			// confirmacao
			confirma_enter()
			clear_screen()

		case "3":
			rand.Seed(time.Now().UnixNano())
			rand.Shuffle(len(array), func(i, j int) { array[i], array[j] = array[j], array[i] })
			fmt.Println("\ndeck embaralhado!")
			// confirmacao
			confirma_enter()
			clear_screen()

		case "4":
			fmt.Println("\nSaindo...")
			return

		default:
			fmt.Println("Opção inválida. Tente novamente.")
		}
	}
}
