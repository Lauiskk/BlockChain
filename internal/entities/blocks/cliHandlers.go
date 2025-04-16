package blocks

import (
	"RedisLike/internal/entities/transaction"
	"flag"
	"fmt"
	"os"
	"strconv"
)

func (cli *CLI) Run(args []string) {
	cli.validateArgs(args)

	addBlockCmd := flag.NewFlagSet("addblock", flag.ExitOnError)
	printChainCmd := flag.NewFlagSet("printchain", flag.ExitOnError)
	getBalanceCmd := flag.NewFlagSet("getbalance", flag.ExitOnError)

	getBalanceAddress := getBalanceCmd.String("address", "", "Endereço para checar o saldo")

	var transactionObj []*transaction.Transaction
	addBlockData := addBlockCmd.String("address", "", "Block data")

	switch args[0] {
	case "addblock":
		err := addBlockCmd.Parse(args[1:])
		if err != nil {
			panic(err)
		}
	case "printchain":
		err := printChainCmd.Parse(args[1:])
		if err != nil {
			panic(err)
		}
	case "getbalance":
		err := getBalanceCmd.Parse(args[1:])
		if err != nil {
			panic(err)
		}
	default:
		cli.printUsage()
		return
	}

	if addBlockCmd.Parsed() {
		if *addBlockData == "" {
			addBlockCmd.Usage()
			return
		}
		cli.addBlock(transactionObj)
	}

	if printChainCmd.Parsed() {
		cli.printChain()
	}

	if getBalanceCmd.Parsed() {
		if *getBalanceAddress == "" {
			getBalanceCmd.Usage()
			return
		}
		cli.getBalance(*getBalanceAddress)
	}
}

func (cli *CLI) validateArgs(args []string) {
	if len(args) < 1 {
		cli.printUsage()
		os.Exit(1)
	}
}

func (cli *CLI) addBlock(transactions []*transaction.Transaction) {
	cli.Bc.AddBlock(transactions)
	fmt.Println("Success!")
}

func (cli *CLI) printChain() {
	bci := cli.Bc.Iterator()

	for {
		block := bci.Next()

		fmt.Printf("Prev. hash: %x\n", block.PrevBlockHash)
		fmt.Printf("Data: %s\n", block.Transactions)
		fmt.Printf("Hash: %x\n", block.Hash)
		pow := NewProofOfWork(block)
		fmt.Printf("PoW: %s\n", strconv.FormatBool(pow.Validate()))
		fmt.Println()

		if len(block.PrevBlockHash) == 0 {
			break
		}
	}
}

func (cli *CLI) printUsage() {
	fmt.Println("Uso:")
	fmt.Println("  addblock -address BLOCK_DATA  -> adiciona um bloco à blockchain")
	fmt.Println("  printchain                 -> imprime todos os blocos da blockchain")
}

func (cli *CLI) getBalance(address string) {
	bc := CreateBlockchain(address)
	defer bc.Db.Close()

	balance := 0
	UTXOs := bc.FindUTXO(address)

	for _, out := range UTXOs {
		balance += out.Value
	}

	fmt.Printf("Balance of '%s': %d\n", address, balance)
}
