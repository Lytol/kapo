package kapo

import "testing"

func TestNewAddressWithValid(t *testing.T) {
	data := [AddressSize]byte{}

	addr, err := NewAddress(data[:])
	if err != nil {
		t.Fatal(err)
	}

	if len(addr) != AddressSize {
		t.Fatal("Incorrect address size")
	}

	for i := 0; i < len(addr); i++ {
		if addr[i] != 0 {
			t.Fatal("Malformed address")
		}
	}
}
