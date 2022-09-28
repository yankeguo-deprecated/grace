package gracex509

import (
	"crypto/rsa"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestPEMPair_Decode(t *testing.T) {
	res, err := Generate(GenerateOptions{Names: []string{"aaa"}, IsCA: true})
	require.NoError(t, err)
	crt, key, err := res.Decode()
	require.NoError(t, err)
	require.Equal(t, "aaa", crt.Subject.CommonName)
	require.IsType(t, &rsa.PrivateKey{}, key)
}
