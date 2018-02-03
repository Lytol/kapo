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

func TestHexToPrivateKeyWithValidKey(t *testing.T) {
	key, err := HexToPrivateKey("ed0f33026b50ab641ade1ac428bfded42b9047c54d5d420d1efc50d7793175f4")

	if err != nil {
		t.Fatal("should not return error")
	}

	if key == nil {
		t.Fatal("should return key")
	}
}

func TestHexToPrivateKeyWithIncorrectLength(t *testing.T) {
	key, err := HexToPrivateKey("ed0f33026b50ab641ade1ac428bfded42b9047c54d5d420d1efc50d7793175")

	if err == nil {
		t.Fatal("should return error")
	}

	if key != nil {
		t.Fatal("should return nil key")
	}
}
