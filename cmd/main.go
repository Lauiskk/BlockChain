package main

import (
	"RedisLike/internal/entities/blocks"
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	cli := blocks.CLI{}
	reader := bufio.NewReader(os.Stdin)

	fmt.Println("Bem-vindo à CLI da Blockchain!")
	fmt.Println("Digite um comando, ou 'exit' para sair.")

	for {
		fmt.Print(">>> ")
		input, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("Erro ao ler entrada:", err)
			continue
		}

		input = strings.TrimSpace(input)
		if input == "exit" {
			break
		}

		args := strings.Fields(input)
		if len(args) == 0 {
			continue
		}

		cli.Run(args)
	}
}
