package kapo

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"crypto/sha256"
	"encoding/hex"
	"math/big"
)

var (
	DefaultCurve = elliptic.P256()
	DefaultRand  = rand.Reader
)

const (
	HashSize       = sha256.Size
	CurveKeySize   = 32 // ECDSA P256 Curve
	PrivateKeySize = CurveKeySize
	PublicKeySize  = CurveKeySize
	SignatureSize  = CurveKeySize
	AddressSize    = sha256.Size
)

//
// Hash
//

type Hash [HashSize]byte

func ToHash(data []byte) Hash {
	var hash Hash

	if len(data) != HashSize {
		panic("Incorrectly sized hash")
	}

	copy(hash[0:HashSize], data)

	return hash
}

func (h Hash) Bytes() []byte { return h[:] }

func HexToHash(s string) Hash {
	decoded, err := hex.DecodeString(s)
	if err != nil {
		panic("Unable to decode hex string")
	}

	return ToHash(decoded)
}

//
// Private Key
//

type PrivateKey struct {
	*ecdsa.PrivateKey
}

func NewPrivateKey() (*PrivateKey, error) {
	key, err := ecdsa.GenerateKey(DefaultCurve, DefaultRand)
	return &PrivateKey{key}, err
}

func (p *PrivateKey) Bytes() []byte {
	return paddedBigBytes(p.D, CurveKeySize)
}

func (p *PrivateKey) Hex() string    { return string(p.Bytes()) }
func (p *PrivateKey) String() string { return p.Hex() }

func (p *PrivateKey) PublicKey() *PublicKey {
	return &PublicKey{p.Public().(*ecdsa.PublicKey)}
}

//
// Public Key
//

type PublicKey struct {
	*ecdsa.PublicKey
}

// Zero-padded, big endian X + Zero-padded, big endian Y
func (p *PublicKey) Bytes() []byte {
	x := paddedBigBytes(p.X, CurveKeySize)
	y := paddedBigBytes(p.Y, CurveKeySize)

	return append(x, y...)
}

func (p *PublicKey) Hex() string    { return string(p.Bytes()) }
func (p *PublicKey) String() string { return p.Hex() }

func (p *PublicKey) Address() Address {
	return PublicKeyToAddress(p)
}

func PublicKeyToAddress(pub *PublicKey) Address {
	hash := SHA(pub.Bytes())
	return ToAddress(hash.Bytes())
}

//
// Address
//

type Address [AddressSize]byte

func ToAddress(data []byte) Address {
	var addr Address

	if len(data) != AddressSize {
		panic("Incorrectly sized hash")
	}

	copy(addr[0:AddressSize], data)

	return addr
}

func (a Address) Bytes() []byte  { return a[:] }
func (a Address) Hex() string    { return string(a.Bytes()) }
func (a Address) String() string { return a.Hex() }

//
// Signature
//

type Signature [SignatureSize]byte

func (s Signature) ToECDSA() (*big.Int, *big.Int) {
	// TODO: concat r + s as 32-byte each
	return nil, nil
}

func (s Signature) Bytes() []byte  { return s[:] }
func (s Signature) Hex() string    { return string(s.Bytes()) }
func (s Signature) String() string { return s.Hex() }

//
// Utility functions
//

func SHA(b []byte) Hash {
	return sha256.Sum256(b)
}

func Sign(priv *PrivateKey, hash Hash) ([]byte, error) {
	// TODO
	return nil, nil
}

func Verify(pub *PublicKey, hash Hash, sig Signature) bool {
	// TODO
	return false
}

const (
	// number of bits in a big.Word
	wordBits = 32 << (uint64(^big.Word(0)) >> 63)
	// number of bytes in a big.Word
	wordBytes = wordBits / 8
)

// paddedBigBytes encodes a big integer as a big-endian byte slice. The length
// of the slice is at least n bytes.
//
// https://github.com/ethereum/go-ethereum/blob/335abdceb1d691e34526f9feb12870d6cdbc3d80/common/math/big.go
//
func paddedBigBytes(num *big.Int, n int) []byte {
	data := make([]byte, n)

	for _, d := range num.Bits() {
		for j := 0; j < wordBytes && n > 0; j++ {
			n--
			data[n] = byte(d)
			d >>= 8
		}
	}

	return data
}
