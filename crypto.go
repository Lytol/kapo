package kapo

import (
	"bytes"
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"crypto/sha256"
	"encoding/hex"
	"errors"
	"fmt"
	"math/big"
)

var (
	DefaultCurve = elliptic.P256()
	DefaultRand  = rand.Reader
)

const (
	HashSize          = sha256.Size
	CurveKeySize      = 32 // ECDSA P256 Curve
	PrivateKeySize    = CurveKeySize
	PublicKeySize     = CurveKeySize * 2
	SignatureHashSize = CurveKeySize * 2
	AddressSize       = sha256.Size
)

//
// Hash
//

var EmptyHash = Hash{}

type Hash [HashSize]byte

func NewHash(data []byte) Hash {
	var hash Hash

	if len(data) != HashSize {
		panic("Incorrectly sized hash")
	}

	copy(hash[0:HashSize], data)

	return hash
}

func (h Hash) Bytes() []byte { return h[:] }
func (h Hash) Hex() string   { return fmt.Sprintf("%x", h) }

func HexToHash(s string) Hash {
	decoded, err := hex.DecodeString(s)
	if err != nil {
		panic("Unable to decode hex string")
	}

	return NewHash(decoded)
}

//
// Private Key
//

type PrivateKey struct {
	key *ecdsa.PrivateKey
}

func NewPrivateKey() (*PrivateKey, error) {
	key, err := ecdsa.GenerateKey(DefaultCurve, DefaultRand)
	return &PrivateKey{key}, err
}

func BytesToPrivateKey(d []byte) (*PrivateKey, error) {
	if len(d) != CurveKeySize {
		return nil, errors.New("Invalid key length")
	}

	priv := new(ecdsa.PrivateKey)
	priv.PublicKey.Curve = DefaultCurve
	priv.D = new(big.Int).SetBytes(d)
	priv.PublicKey.X, priv.PublicKey.Y = priv.PublicKey.Curve.ScalarBaseMult(d)

	if priv.PublicKey.X == nil || priv.PublicKey.Y == nil {
		return nil, errors.New("invalid private key")
	}

	return &PrivateKey{priv}, nil
}

func HexToPrivateKey(str string) (*PrivateKey, error) {
	b, err := HexToBytes(str)
	if err != nil {
		return nil, err
	}

	return BytesToPrivateKey(b)
}

func (p *PrivateKey) Bytes() []byte {
	return paddedBigBytes(p.key.D, CurveKeySize)
}

func (p *PrivateKey) Hex() string    { return fmt.Sprintf("%x", p.Bytes()) }
func (p *PrivateKey) String() string { return p.Hex() }

func (p *PrivateKey) PublicKey() *PublicKey {
	return &PublicKey{p.key.Public().(*ecdsa.PublicKey)}
}

//
// Public Key
//

type PublicKey struct {
	key *ecdsa.PublicKey
}

// Zero-padded, big endian X + Zero-padded, big endian Y
func (p *PublicKey) Bytes() []byte {
	x := paddedBigBytes(p.key.X, CurveKeySize)
	y := paddedBigBytes(p.key.Y, CurveKeySize)

	return append(x, y...)
}

func (p *PublicKey) Hex() string    { return fmt.Sprintf("%x", p.Bytes()) }
func (p *PublicKey) String() string { return p.Hex() }

func (p *PublicKey) Address() Address {
	return PublicKeyToAddress(p)
}

func PublicKeyToAddress(pub *PublicKey) Address {
	hash := SHA(pub.Bytes())
	addr, _ := NewAddress(hash.Bytes())
	return addr
}

//
// Address
//

type Address [AddressSize]byte

func NewAddress(data []byte) (Address, error) {
	var addr Address

	if len(data) != AddressSize {
		return addr, errors.New("Incorrectly sized hash")
	}

	copy(addr[0:AddressSize], data)

	return addr, nil
}

func HexToAddress(str string) (Address, error) {
	b, err := HexToBytes(str)
	if err != nil {
		return Address{}, err
	}

	return NewAddress(b)
}

func (a Address) Bytes() []byte  { return a[:] }
func (a Address) Hex() string    { return fmt.Sprintf("%x", a.Bytes()) }
func (a Address) String() string { return a.Hex() }

//
// Signature
//

type Signature struct {
	R         *big.Int
	S         *big.Int
	PublicKey *PublicKey
}

func (s *Signature) Bytes() []byte {
	var buf bytes.Buffer

	buf.Write(paddedBigBytes(s.R, CurveKeySize))
	buf.Write(paddedBigBytes(s.S, CurveKeySize))
	buf.Write(s.PublicKey.Bytes())

	return buf.Bytes()
}

func (s *Signature) Hex() string    { return fmt.Sprintf("%x", s.Bytes()) }
func (s *Signature) String() string { return s.Hex() }

//
// Utility functions
//

func SHA(b []byte) Hash {
	return sha256.Sum256(b)
}

func Sign(hash Hash, priv *PrivateKey) (*Signature, error) {
	r, s, err := ecdsa.Sign(DefaultRand, priv.key, hash.Bytes())
	if err != nil {
		return nil, err
	}

	return &Signature{r, s, priv.PublicKey()}, nil
}

func Verify(hash Hash, sig *Signature) bool {
	return ecdsa.Verify(sig.PublicKey.key, hash.Bytes(), sig.R, sig.S)
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
