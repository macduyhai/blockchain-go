package proofofwork

import (
	"bytes"
	"crypto/sha256"
	"fmt"
	b "go-block-chain/block"
	"go-block-chain/ultis"
	"math"
	"math/big"
	"time"
)

var (
	maxNonce = math.MaxInt64
)

const targetBits = 24

// ProofOfWork represents a proof-of-work
type ProofOfWork struct {
	block  *b.Block
	target *big.Int
}

// NewBlock creates and returns Block
func NewBlock(data string, prevBlockHash []byte) *b.Block {
	block := &b.Block{time.Now().Unix(), []byte(data), prevBlockHash, []byte{}, 0}
	pow := NewProofOfWork(block)
	nonce, hash := pow.Run()

	block.Hash = hash[:]
	block.Nonce = nonce

	return block
}

// NewGenesisBlock creates and returns genesis Block
func NewGenesisBlock() *b.Block {
	return NewBlock("Genesis Block", []byte{})
}

// NewProofOfWork builds and returns a ProofOfWork
func NewProofOfWork(b *b.Block) *ProofOfWork {
	target := big.NewInt(1)
	target.Lsh(target, uint(256-targetBits))

	pow := &ProofOfWork{b, target}

	return pow
}

func (pow *ProofOfWork) prepareData(nonce int) []byte {
	data := bytes.Join(
		[][]byte{
			pow.block.PrevBlockHash,
			pow.block.Data,
			ultis.IntToHex(pow.block.Timestamp),
			ultis.IntToHex(int64(targetBits)),
			ultis.IntToHex(int64(nonce)),
		},
		[]byte{},
	)

	return data
}

// Run performs a proof-of-work
func (pow *ProofOfWork) Run() (int, []byte) {
	var hashInt big.Int
	var hash [32]byte
	nonce := 0

	fmt.Printf("Mining the block containing \"%s\"\n", pow.block.Data)
	for nonce < maxNonce {
		data := pow.prepareData(nonce)

		hash = sha256.Sum256(data)
		fmt.Printf("\r%x", hash)
		hashInt.SetBytes(hash[:])

		if hashInt.Cmp(pow.target) == -1 {
			break
		} else {
			nonce++
		}
	}
	fmt.Print("\n\n")

	return nonce, hash[:]
}

// Validate validates block's PoW
func (pow *ProofOfWork) Validate() bool {
	var hashInt big.Int

	data := pow.prepareData(pow.block.Nonce)
	hash := sha256.Sum256(data)
	hashInt.SetBytes(hash[:])

	isValid := hashInt.Cmp(pow.target) == -1

	return isValid
}
