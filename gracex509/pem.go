package gracex509

import (
	"crypto/x509"
	"encoding/pem"
	"errors"
	"github.com/guoyk93/grace/gracepem"
)

const (
	PEMTypeCertificate   = "CERTIFICATE"
	PEMTypeRSAPrivateKey = "RSA PRIVATE KEY"
	PEMTypePrivateKey    = "PRIVATE KEY"
)

// PEMPair PEM encoded x509 key pair
type PEMPair struct {
	Crt []byte
	Key []byte
}

func (b PEMPair) IsZero() bool {
	return len(b.Crt)+len(b.Key) == 0
}

func (b PEMPair) Certificate() (crt *x509.Certificate, err error) {
	var buf []byte
	if buf, err = gracepem.DecodeSingle(b.Crt, PEMTypeCertificate); err != nil {
		return
	}
	crt, err = x509.ParseCertificate(buf)
	return
}

func (b PEMPair) PrivateKey() (key any, err error) {
	var p *pem.Block
	if p, _ = pem.Decode(b.Key); p == nil {
		err = errors.New("PEMPair.Key: missing PEM block")
		return
	}
	switch p.Type {
	case PEMTypePrivateKey:
		key, err = x509.ParsePKCS8PrivateKey(p.Bytes)
	case PEMTypeRSAPrivateKey:
		key, err = x509.ParsePKCS1PrivateKey(p.Bytes)
	default:
		err = errors.New("PEMPair.Key: unknown PEM block type: " + p.Type)
	}
	return
}

func (b PEMPair) Decode() (crt *x509.Certificate, key any, err error) {
	if crt, err = b.Certificate(); err != nil {
		return
	}
	if key, err = b.PrivateKey(); err != nil {
		return
	}
	return
}
