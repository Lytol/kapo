package kapo

import (
	"bytes"
	"crypto/sha256"
	"encoding/binary"
)

const HashSize = sha256.Size

type Hash = [HashSize]byte

func Int64ToBytes(n int64) []byte {
	buf := new(bytes.Buffer)

	err := binary.Write(buf, binary.BigEndian, n)
	if err != nil {
		panic(err)
	}

	return buf.Bytes()
}
