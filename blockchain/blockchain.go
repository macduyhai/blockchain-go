package blockchain

import (
	b "go-block-chain/block"
)

type BlockChain struct {
	Blocks []*b.Block
}

func NewGenesisBlock() *b.Block {
	return b.NewBlock("Genesis Block", []byte{})
}
func NewBlockChain() *BlockChain {
	return &BlockChain{[]*b.Block{NewGenesisBlock()}}
}
func (bc *BlockChain) AddBlock(data string) {
	prevBlock := bc.Blocks[len(bc.Blocks)-1]
	newBlock := b.NewBlock(data, prevBlock.Hash)
	bc.Blocks = append(bc.Blocks, newBlock)
}
