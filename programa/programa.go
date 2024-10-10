// go mod init programa
// go get -u github.com/nsf/termbox-go  

package main

import (
    "fmt"
    "time"
    "github.com/nsf/termbox-go"
)

func draw() {
    // Limpa a tela
    termbox.Clear(termbox.ColorDefault, termbox.ColorDefault)

    // Desenha alguma coisa na tela (neste caso, apenas uma string)
    msg := "Pressione 'q' para sair"
    for i, c := range msg {
        termbox.SetCell(i, 0, c, termbox.ColorDefault, termbox.ColorDefault)
    }

    // Atualiza a tela
    termbox.Flush()
}

func main() {
    // Inicializa a biblioteca termbox
    err := termbox.Init()
    if err != nil {
        fmt.Println("Erro ao inicializar termbox:", err)
        return
    }
    defer termbox.Close()

    // Desenha a tela inicial
    draw()

    // Loop principal
    ticker := time.NewTicker(time.Second)
    defer ticker.Stop()

    // Canal para sinais de saída
    quit := make(chan struct{})

    go func() {
        for {
            switch ev := termbox.PollEvent(); ev.Type {
            case termbox.EventKey:
                if ev.Ch == 'q' {
                    close(quit)
                    return
                }
            }
        }
    }()

    for {
        select {
        case <-ticker.C:
            // Atualiza a tela a cada segundo
            draw()
        case <-quit:
            // Sai do programa quando 'q' for pressionado
            return
        }
    }
}
