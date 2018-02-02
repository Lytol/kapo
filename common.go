package kapo

import (
	"bytes"
	"encoding/binary"
	"encoding/hex"
)

//
// Utility functions
//

func Int64ToBytes(n int64) []byte {
	buf := new(bytes.Buffer)

	err := binary.Write(buf, binary.BigEndian, n)
	if err != nil {
		panic(err)
	}

	return buf.Bytes()
}

func HexToBytes(s string) ([]byte, error) {
	if len(s) > 1 {
		if s[0:2] == "0x" || s[0:2] == "0X" {
			s = s[2:]
		}
	}

	if len(s)%2 == 1 {
		s = "0" + s
	}

	decoded, err := hex.DecodeString(s)
	if err != nil {
		return []byte{}, err
	}

	return decoded, nil
}
