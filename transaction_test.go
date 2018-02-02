package kapo

import (
	"testing"
)

func TestNewTransactionWithValid(t *testing.T) {
	addr, _ := HexToAddress("7c08b268f322f39e8e697a4f86db3365f78fc094910cec02d4e9d921c2e00955")
	data := []byte("This is a test")

	tx := NewTransaction(addr, data)

	if tx == nil {
		t.Fatal("Transaction must not be nil")
	}

	if tx.Address != addr {
		t.Fatal("Address should be set")
	}

	if string(tx.Data) != "This is a test" {
		t.Fatal("Data should be set")
	}

	if tx.Hash != EmptyHash {
		t.Fatal("Hash should be empty")
	}

	if tx.Signature != nil {
		t.Fatal("Signature should be nil")
	}
}

func TestTransactionHasSignatureWithoutSignature(t *testing.T) {
	addr, _ := HexToAddress("7c08b268f322f39e8e697a4f86db3365f78fc094910cec02d4e9d921c2e00955")

	tx := NewTransaction(addr, []byte("testing"))

	if tx.HasSignature() {
		t.Fatal("Transaction should not be signed")
	}
}

func TestTransactionHasSignatureWithValidSignature(t *testing.T) {
	account, _ := NewAccount()

	tx := NewTransaction(account.Address, []byte("testing"))
	tx.Sign(account.PrivateKey)

	if tx.HasSignature() == false {
		t.Fatal("Transaction should be signed")
	}

	t.Log(tx)
}
