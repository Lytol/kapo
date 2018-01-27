package kapo

import (
	"bytes"
	"crypto/sha256"
	"encoding/binary"
	"encoding/hex"
)

const (
	HashSize = sha256.Size
)

type Hash [HashSize]byte

func NewHash(data []byte) Hash {
	var hash Hash

	if len(data) != HashSize {
		panic("Incorrectly sized hash")
	}

	copy(hash[0:HashSize], data)

	return hash
}

func (h Hash) Bytes() []byte {
	return h[:]
}

func Int64ToBytes(n int64) []byte {
	buf := new(bytes.Buffer)

	err := binary.Write(buf, binary.BigEndian, n)
	if err != nil {
		panic(err)
	}

	return buf.Bytes()
}

func HexStringToHash(s string) Hash {
	decoded, err := hex.DecodeString(s)
	if err != nil {
		panic("Unable to decode hex string")
	}

	return NewHash(decoded)
}
