package kapo

import (
	"testing"
)

func TestHexToBytesWithNormalString(t *testing.T) {
	hex := "ff"

	data, err := HexToBytes(hex)

	if err != nil {
		t.Fatal(err)
	}

	if len(data) != 1 && data[0] != 255 {
		t.Fatalf("Expected 255, received %v", data[0])
	}
}
