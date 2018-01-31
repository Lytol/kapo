package kapo

import (
	"bytes"
	"fmt"
)

type Transaction struct {
	Hash      Hash
	Address   Address
	Script    []byte
	Data      []byte
	Signature Signature
}

func NewTransaction(addr Address, script []byte, data []byte) *Transaction {
	tx := &Transaction{
		Address: addr,
		Script:  script,
		Data:    data,
	}

	tx.Hash = tx.getHash()

	return tx
}

func (tx *Transaction) Sign(priv *PrivateKey) error {
	return nil
}

func (tx *Transaction) getHash() Hash {
	var buf bytes.Buffer

	// TODO: No error checking
	buf.Write(tx.Address.Bytes())
	buf.Write(tx.Script)
	buf.Write(tx.Data)

	return SHA(buf.Bytes())
}

func (tx *Transaction) String() string {
	return fmt.Sprintf(`
	Transaction(%x)
	Address: %x
	Signature: %x
	Script: %s
	Data: %s
`, tx.Hash, tx.Address.Hex(), tx.Signature.Hex(), string(tx.Script), string(tx.Data))
}
