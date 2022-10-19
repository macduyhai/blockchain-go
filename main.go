package main

import (
	bc "go-block-chain/blockchain"
	"go-block-chain/proof"
	"log"
	"strconv"
)

func main() {
	blockChain := bc.NewBlockChain()
	blockChain.AddBlock("Simon")
	blockChain.AddBlock("Mac Duy Hai")
	for _, block := range blockChain.Blocks {
		log.Printf("Prev block's hash: %x\n", block.PrevBlockHash)
		log.Printf("Current block's hash: %x\n", block.Hash)
		log.Printf("Current block's data: %x\n", block.Data)
		pow := proof.NewProofOfWork(block)
		log.Printf("PoW: %s\n", strconv.FormatBool(pow.Validate()))
		log.Println()
	}
}
