package kapo

import (
	"bytes"
	"crypto/sha256"
	"encoding/gob"
	"log"
)

type Transaction struct {
	ID        Hash
	Recipient string
	Signature Hash
	Data      []byte
}

func NewTransaction(recipient string, data []byte) *Transaction {
	tx := &Transaction{
		Recipient: recipient,
		Data:      data,
	}

	tx.ID = tx.hash()

	return tx
}

func (tx *Transaction) hash() Hash {
	var encoded bytes.Buffer

	enc := gob.NewEncoder(&encoded)

	err := enc.Encode(tx)
	if err != nil {
		log.Panic(err)
	}

	return sha256.Sum256(encoded.Bytes())
}
