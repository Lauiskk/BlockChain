package blocks

import (
	"flag"
	"fmt"
	"os"
	"strconv"
)

func (cli *CLI) Run(args []string) {
	cli.validateArgs(args)

	addBlockCmd := flag.NewFlagSet("addblock", flag.ExitOnError)
	printChainCmd := flag.NewFlagSet("printchain", flag.ExitOnError)

	addBlockData := addBlockCmd.String("data", "", "Block data")

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
	default:
		cli.printUsage()
		return
	}

	if addBlockCmd.Parsed() {
		if *addBlockData == "" {
			addBlockCmd.Usage()
			return
		}
		cli.addBlock(*addBlockData)
	}

	if printChainCmd.Parsed() {
		cli.printChain()
	}
}

func (cli *CLI) validateArgs(args []string) {
	if len(args) < 1 {
		cli.printUsage()
		os.Exit(1)
	}
}

func (cli *CLI) addBlock(data string) {
	cli.Bc.AddBlock(data)
	fmt.Println("Success!")
}

func (cli *CLI) printChain() {
	bci := cli.Bc.Iterator()

	for {
		block := bci.Next()

		fmt.Printf("Prev. hash: %x\n", block.PrevBlockHash)
		fmt.Printf("Data: %s\n", block.Data)
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
	fmt.Println("  addblock -data BLOCK_DATA  -> adiciona um bloco à blockchain")
	fmt.Println("  printchain                 -> imprime todos os blocos da blockchain")
}
