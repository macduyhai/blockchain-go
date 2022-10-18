package main

import (
	bc "go-block-chain/blockchain"
	"log"
)

func main() {
	blockChain := bc.NewBlockChain()
	blockChain.AddBlock("Simon")
	blockChain.AddBlock("Mac Duy Hai")
	for _, block := range blockChain.Blocks {
		log.Printf("Prev block's hash: %x\n", block.PrevBlockHash)
		log.Printf("Current block's hash: %x\n", block.Hash)
		log.Printf("Current block's data: %x\n", block.Data)
		log.Println()
	}
}
