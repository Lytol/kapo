package kapo

import (
	"fmt"
)

type Account struct {
	PrivateKey *PrivateKey
	PublicKey  *PublicKey
	Address    Address
}

func NewAccount() (*Account, error) {
	priv, err := NewPrivateKey()
	if err != nil {
		return nil, err
	}

	pub := priv.PublicKey()
	addr := pub.Address()

	return &Account{priv, pub, addr}, nil
}

func (a *Account) String() string {
	return fmt.Sprintf(`Account
  Address:    %s
  PrivateKey: %s
  PublicKey:  %s`,
		a.Address.Hex(), a.PrivateKey.Hex(), a.PublicKey.Hex())
}
