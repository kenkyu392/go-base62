package base62

import (
	"math"
	"math/big"
)

const (
	// GMP: 0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz
	// GMP INVERTED: 0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ
	encodeStd = "0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz"
	baseLen   = int64(len(encodeStd))
)

var baseLog = math.Log(float64(baseLen)*4) / math.Log(float64(baseLen))

// StdEncoding is the standard base62 encoding.
var StdEncoding = NewEncoding(encodeStd)

// Encoding ...
type Encoding struct {
	encode [62]byte
}

// NewEncoding returns a new Encoding defined by the given alphabet,
// which must be a 62-byte string.
func NewEncoding(encoder string) *Encoding {
	if len(encoder) != 62 {
		panic("encoding alphabet is not 62-bytes long")
	}
	for i := 0; i < len(encoder); i++ {
		if encoder[i] == '\n' || encoder[i] == '\r' {
			panic("encoding alphabet contains newline character")
		}
	}
	e := new(Encoding)
	copy(e.encode[:], encoder)
	return e
}

// Encode encodes src using the encoding enc, writing
// EncodedLen(len(src)) bytes to dst.
func (enc *Encoding) Encode(dst, src []byte) {
	if len(src) == 0 {
		return
	}
	// enc is a pointer receiver, so the use of enc.encode within the hot
	// loop below means a nil check at every operation. Lift that nil check
	// outside of the loop to speed up the encoder.
	_ = enc.encode
	var (
		i    = new(big.Int).SetBytes(src)
		base = big.NewInt(baseLen)
		zero = big.NewInt(0)
		rem  = big.NewInt(0)
	)
	for n := len(dst) - 1; i.Cmp(zero) == 1; n-- {
		i, rem = i.DivMod(i, base, rem)
		dst[n] = enc.encode[rem.Int64()]
	}
}

// EncodeToString returns the base62 encoding of src.
func (enc *Encoding) EncodeToString(src []byte) string {
	buf := make([]byte, enc.EncodedLen(len(src)))
	enc.Encode(buf, src)
	return string(buf)
}

// EncodedLen returns the length in bytes of the base62 encoding
// of an input buffer of length n.
func (enc *Encoding) EncodedLen(n int) int {
	return int(math.Round(float64(n) * baseLog))
}
