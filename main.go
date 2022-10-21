package main

import (
	"go-block-chain/blockchain"
	"go-block-chain/cli"
)

func main() {
	bc := blockchain.NewBlockchain()
	defer bc.DB.Close()

	cli := cli.CLI{bc}
	cli.Run()
}
