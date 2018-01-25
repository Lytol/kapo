package kapo

import (
	"bytes"
	"crypto/sha256"
	"encoding/binary"
	"errors"
)

const (
	HashSize = sha256.Size
)

type Hash [HashSize]byte

func NewHash(data []byte) (Hash, error) {
	var hash Hash

	if len(data) != HashSize {
		return Hash{}, errors.New("Incorrectly sized hash")
	}

	copy(hash[0:HashSize], data)

	return hash, nil
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
