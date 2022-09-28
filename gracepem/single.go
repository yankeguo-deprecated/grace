package gracepem

import (
	"encoding/pem"
	"errors"
)

// DecodeSingle decode a single block from PEM encoded data
// if typ is not an empty string, will search for specified type
func DecodeSingle(buf []byte, typ string) (out []byte, err error) {
	var b *pem.Block
	for {
		b, buf = pem.Decode(buf)

		if b == nil {
			if typ == "" {
				err = errors.New("missing PEM block")
			} else {
				err = errors.New("missing PEM block with type: " + typ)
			}
			return
		} else {
			if typ == "" || typ == b.Type {
				out = b.Bytes
				return
			}
		}
	}
}

// EncodeSingle encode single block as PEM format
func EncodeSingle(buf []byte, typ string) []byte {
	return pem.EncodeToMemory(&pem.Block{
		Type:  typ,
		Bytes: buf,
	})
}
