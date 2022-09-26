package gracex509

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"crypto/x509/pkix"
	"encoding/pem"
	"errors"
	"math/big"
	"time"
)

func createCertificate(template, parent *x509.Certificate, pub any, priv any) (crtOut *x509.Certificate, crtPEMOut []byte, err error) {
	var crtRaw []byte
	if crtRaw, err = x509.CreateCertificate(rand.Reader, template, parent, pub, priv); err != nil {
		return
	}

	if crtOut, err = x509.ParseCertificate(crtRaw); err != nil {
		return
	}

	crtPEMOut = pem.EncodeToMemory(&pem.Block{
		Type:  "CERTIFICATE",
		Bytes: crtRaw,
	})

	return
}

const (
	DefaultCountry      = "CN"
	DefaultOrganization = "guoyk93.github.io"
	DefaultExpires      = time.Hour * 24 * 365 * 30
)

// GenerationOptions options for certificate generation
type GenerationOptions struct {
	// CACrtPEM public key pem for ca, required for leaf certificate generation
	CACrtPEM []byte
	// CAKeyPEM private key pem for ca, required for leaf certificate generation
	CAKeyPEM []byte
	// Names certificate names, tailing names will be used as DNSNames
	Names []string
	// Country certificate country
	Country string
	// Organization certificate organization
	Organization string
	// Expires certificate duration
	Expires time.Duration
}

func (opts GenerationOptions) IsCA() bool {
	return len(opts.CACrtPEM)+len(opts.CAKeyPEM) == 0
}

// GenerationResult result of certificate generation
type GenerationResult struct {
	Crt    *x509.Certificate
	CrtPEM []byte
	Key    *rsa.PrivateKey
	KeyPEM []byte
}

// Generate easily generate certificate and private key
// if both CACrtPEM and CAKeyPem is missing, will generate a new CA
func Generate(opts GenerationOptions) (res GenerationResult, err error) {
	// check options
	if len(opts.Names) < 1 {
		err = errors.New("gracex509.Generate: opts.Names missing")
		return
	}
	if opts.Country == "" {
		opts.Country = DefaultCountry
	}
	if opts.Organization == "" {
		opts.Organization = DefaultOrganization
	}
	if opts.Expires <= 0 {
		opts.Expires = DefaultExpires
	}

	// generate res.Key
	if res.Key, err = rsa.GenerateKey(rand.Reader, 2048); err != nil {
		return
	}
	res.KeyPEM = pem.EncodeToMemory(&pem.Block{
		Type:  "RSA PRIVATE KEY",
		Bytes: x509.MarshalPKCS1PrivateKey(res.Key),
	})

	var (
		notBefore = time.Now().Add(-time.Second * 10)
		notAfter  = notBefore.Add(opts.Expires)
	)

	var template *x509.Certificate

	if opts.IsCA() {
		// ca certificate
		template = &x509.Certificate{
			SerialNumber: big.NewInt(notBefore.UnixMilli()),
			Subject: pkix.Name{
				Country:      []string{opts.Country},
				Organization: []string{opts.Organization},
				CommonName:   opts.Names[0],
			},
			NotBefore:             notBefore,
			NotAfter:              notAfter,
			KeyUsage:              x509.KeyUsageCertSign | x509.KeyUsageCRLSign,
			ExtKeyUsage:           []x509.ExtKeyUsage{x509.ExtKeyUsageServerAuth},
			BasicConstraintsValid: true,
			IsCA:                  true,
			MaxPathLen:            2,
		}

		// assign res.Crt res.CrtPEM
		if res.Crt, res.CrtPEM, err = createCertificate(
			template,
			template,
			&res.Key.PublicKey,
			res.Key,
		); err != nil {
			return
		}
	} else {
		// leaf certificate
		var (
			caCrt *x509.Certificate
			caKey *rsa.PrivateKey
		)
		{
			caCrtRaw, _ := pem.Decode(opts.CACrtPEM)
			if caCrtRaw == nil {
				err = errors.New("gracex509.Generate: invalid opts.CACrtPEM")
				return
			}
			if caCrtRaw.Type != "CERTIFICATE" {
				err = errors.New("gracex509.Generate: invalid opts.CACrtPEM pem type: " + caCrtRaw.Type)
				return
			}
			if caCrt, err = x509.ParseCertificate(caCrtRaw.Bytes); err != nil {
				return
			}
		}
		{
			caKeyRaw, _ := pem.Decode(opts.CAKeyPEM)
			if caKeyRaw == nil {
				err = errors.New("gracex509.Generate: invalid opts.CAKeyPEM")
				return
			}
			if caKeyRaw.Type != "RSA PRIVATE KEY" {
				err = errors.New("gracex509.Generate: invalid opts.CAKeyPEM pem type: " + caKeyRaw.Type)
				return
			}
			if caKey, err = x509.ParsePKCS1PrivateKey(caKeyRaw.Bytes); err != nil {
				return
			}
		}
		template = &x509.Certificate{
			SerialNumber: big.NewInt(notBefore.UnixMilli()),
			Subject: pkix.Name{
				Country:      []string{opts.Country},
				Organization: []string{opts.Organization},
				CommonName:   opts.Names[0],
			},
			DNSNames:    opts.Names[1:],
			NotBefore:   notBefore,
			NotAfter:    notAfter,
			KeyUsage:    x509.KeyUsageDigitalSignature | x509.KeyUsageKeyEncipherment,
			ExtKeyUsage: []x509.ExtKeyUsage{x509.ExtKeyUsageServerAuth},
		}
		if res.Crt, res.CrtPEM, err = createCertificate(template, caCrt, &res.Key.PublicKey, caKey); err != nil {
			return
		}
	}

	return
}
