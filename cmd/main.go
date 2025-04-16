package main

import (
	"RedisLike/internal/entities/blocks"
	"bufio"
	"fmt"
	"github.com/boltdb/bolt"
	"os"
	"strings"
)

func main() {
	bc := blocks.NewBlockchain()

	if bc.Db == nil {
		fmt.Println("Erro: o banco de dados não foi aberto corretamente!")
		os.Exit(1)
	}

	defer func(Db *bolt.DB) {
		err := Db.Close()
		if err != nil {
			fmt.Println("Erro ao fechar o banco:", err)
		}
	}(bc.Db)

	cli := blocks.CLI{Bc: bc}
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
