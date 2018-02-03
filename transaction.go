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

	hash := tx.signatureHash()

	tx.Signature, err = Sign(hash, priv)
	if err != nil {
		return err
	}

	tx.Hash = tx.getHash()

	return nil
}

func (tx *Transaction) Verify() bool {
	if tx.HasSignature() == false {
		return false
	}

	signatureAddress := PublicKeyToAddress(tx.Signature.PublicKey)
	if tx.Address != signatureAddress {
		return false
	}

	return Verify(tx.signatureHash(), tx.Signature)
}

func (tx *Transaction) HasSignature() bool {
	return tx.Hash != EmptyHash && tx.Signature != nil
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
  Hash:      %s
  Address:   %s
  Data:      %s`,
		tx.Hash.Hex(),
		tx.Address.Hex(),
		string(tx.Data))

	if tx.HasSignature() {
		str = str + fmt.Sprintf("\n  Signature: %s", tx.Signature.Hex())
		str = str + "\n  Signed:    ✔"
	} else {
		str = str + "\n  Signed:    ✖"
	}

	if tx.Verify() {
		str = str + "\n  Verified:  ✔"
	} else {
		str = str + "\n  Verified:  ✖"
	}

	return str
}

func (tx *Transaction) encodeForSignature() []byte {
	var buf bytes.Buffer

	// TODO: No error checking, extract method
	buf.Write(tx.Address.Bytes())
	buf.Write(tx.Data)

	return buf.Bytes()
}

func (tx *Transaction) signatureHash() Hash {
	return SHA(tx.encodeForSignature())
}
