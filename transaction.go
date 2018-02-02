package kapo

import (
	"bytes"
	"fmt"
)

type Transaction struct {
	Hash      Hash
	Address   Address
	Data      []byte
	Signature *Signature
}

func NewTransaction(addr Address, data []byte) *Transaction {
	tx := &Transaction{
		Address: addr,
		Data:    data,
	}

	return tx
}

func (tx *Transaction) Sign(priv *PrivateKey) error {
	var err error
	var buf bytes.Buffer

	// TODO: No error checking, extract method
	buf.Write(tx.Address.Bytes())
	buf.Write(tx.Data)

	h := SHA(buf.Bytes())

	tx.Signature, err = Sign(h, priv)
	if err != nil {
		return err
	}

	tx.Hash = tx.getHash()

	return nil
}

func (tx *Transaction) getHash() Hash {
	var buf bytes.Buffer

	// TODO: No error checking, extract method
	buf.Write(tx.Address.Bytes())
	buf.Write(tx.Data)
	buf.Write(tx.Signature.Bytes())

	return SHA(buf.Bytes())
}

func (tx *Transaction) String() string {
	str := fmt.Sprintf(`Transaction
  Hash:    %s
  Address: %s
  Data:    %s`,
		tx.Hash.Hex(),
		tx.Address.Hex(),
		string(tx.Data))

	if tx.Signature != nil {
		str = str + "\n  Signed:  ✔"
	} else {
		str = str + "\n  Signed:  ✖"
	}

	// TODO: Show verified yes/no
	// TODO: Show signature if signed

	return str
}
