package kapo

import (
	"log"
	"time"
)

type Block struct {
	ID           Hash
	Previous     Hash
	Timestamp    int64
	Transactions []*Transaction
	Nonce        int64
}

func NewBlock(txs []*Transaction, previous Hash) *Block {
	block := &Block{
		Previous:     previous,
		Timestamp:    time.Now().Unix(),
		Transactions: txs,
	}

	engine := &ProofOfWork{}

	block, err := engine.Run(block)
	if err != nil {
		log.Fatal(err)
	}

	return block
}

var GenesisBlock = NewBlock([]*Transaction{}, Hash{})
