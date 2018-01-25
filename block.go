package kapo

import (
	"bytes"
	"encoding/gob"
	"log"
	"time"
)

var GenesisBlock = NewBlock([]*Transaction{}, Hash{})

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

func (b *Block) Serialize() ([]byte, error) {
	var buf bytes.Buffer

	encoder := gob.NewEncoder(&buf)

	err := encoder.Encode(b)
	if err != nil {
		return []byte{}, err
	}

	return buf.Bytes(), nil
}

func (b *Block) Deserialize(data []byte) error {
	decoder := gob.NewDecoder(bytes.NewReader(data))

	err := decoder.Decode(b)
	if err != nil {
		return err
	}

	return nil
}
