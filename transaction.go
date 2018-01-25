package kapo

import (
	"bytes"
	"crypto/sha256"
	"fmt"
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
	var buf bytes.Buffer

	// TODO: No error checking
	buf.Write(tx.ID.Bytes())
	buf.WriteString(tx.Recipient)
	buf.Write(tx.Signature.Bytes())
	buf.Write(tx.Data)

	return sha256.Sum256(buf.Bytes())
}

func (tx *Transaction) String() string {
	return fmt.Sprintf("%x | %s | %s", tx.ID, tx.Recipient, string(tx.Data))
}
