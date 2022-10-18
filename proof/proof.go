package proof

import (
	"bytes"
	"crypto/sha256"
	"encoding/binary"
	"fmt"
	"go-block-chain/block"
	"math"
	"math/big"
)

const targetBit = 24

type ProofOfWork struct {
	block  *block.Block
	target *big.Int
}

func (pow *ProofOfWork) Run() (int, []byte) {
	var hashInt big.Int
	var hash [32]byte
	nonce := 0
	maxNonce := math.MaxInt64

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

func NewProofOfWork(b *block.Block) *ProofOfWork {
	target := big.NewInt(1)
	target.Lsh(target, uint(256-targetBit))
	pow := &ProofOfWork{b, target}
	return pow
}
func (pow *ProofOfWork) prepareData(nonce int) []byte {

	data := bytes.Join(
		[][]byte{
			pow.block.PrevBlockHash,
			pow.block.Data,
			IntToHex(pow.block.Timestamp),
			IntToHex(int64(targetBit)),
			IntToHex(int64(nonce)),
		},
		[]byte{},
	)
	return data
}

// strconv.ParseInt(pow.block.Timestamp, 16, 64),
func IntToHex(num int64) []byte {
	b := make([]byte, 8)
	binary.LittleEndian.PutUint64(b, uint64(num))
	return b
}
