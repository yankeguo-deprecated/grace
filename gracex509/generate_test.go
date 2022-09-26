package gracex509

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestGenerate(t *testing.T) {
	resCA, err := Generate(GenerationOptions{
		Names: []string{"test-ca"},
	})
	require.NoError(t, err)
	require.True(t, resCA.Crt.BasicConstraintsValid)
	require.True(t, resCA.Crt.IsCA)
	resUR, err := Generate(GenerationOptions{
		CACrtPEM: resCA.CrtPEM,
		CAKeyPEM: resCA.KeyPEM,
		Names:    []string{"test-leaf"},
	})
	require.NoError(t, err)
	require.True(t, resCA.Crt.BasicConstraintsValid)
	require.False(t, resUR.Crt.IsCA)
}
