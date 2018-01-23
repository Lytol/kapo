package kapo

import (
	"bytes"
	"crypto/sha256"
	"math"
	"math/big"
)

type Engine interface {
	Run(block *Block) (*Block, error)
}

const targetBits = 20

type ProofOfWork struct{}

func (pow *ProofOfWork) Run(block *Block) (*Block, error) {
	var hash Hash
	var hashInt big.Int

	target := big.NewInt(1)
	target.Lsh(target, uint(256-targetBits))

	var nonce int64 = 0

	for nonce < math.MaxInt64 {
		hash = pow.hash(block, nonce)
		hashInt.SetBytes(hash[:])

		if hashInt.Cmp(target) == -1 {
			break
		} else {
			nonce++
		}
	}

	block.Nonce = nonce
	block.ID = hash

	return block, nil
}

func (pow *ProofOfWork) hash(block *Block, nonce int64) Hash {
	txsHash := pow.transactionsHash(block)

	headers := bytes.Join([][]byte{
		block.Previous[:],
		txsHash[:],
		Int64ToBytes(block.Timestamp),
		Int64ToBytes(nonce),
	}, []byte{})

	return sha256.Sum256(headers)
}

func (pow *ProofOfWork) transactionsHash(block *Block) Hash {
	var txHashes [][]byte

	for _, tx := range block.Transactions {
		txHashes = append(txHashes, tx.ID[:])
	}

	return sha256.Sum256(bytes.Join(txHashes, []byte{}))
}
