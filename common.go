package kapo

import (
	"bytes"
	"encoding/binary"
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
