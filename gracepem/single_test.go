package gracepem

import (
	"bytes"
	"encoding/pem"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestDecodeSingle(t *testing.T) {
	buf := &bytes.Buffer{}

	err := pem.Encode(buf, &pem.Block{
		Type:  "AAA",
		Bytes: []byte("aaa"),
	})
	require.NoError(t, err)

	err = pem.Encode(buf, &pem.Block{
		Type:  "BBB",
		Bytes: []byte("bbb"),
	})
	require.NoError(t, err)

	out, err := DecodeSingle(buf.Bytes(), "")
	require.NoError(t, err)
	require.Equal(t, []byte("aaa"), out)

	out, err = DecodeSingle(buf.Bytes(), "BBB")
	require.NoError(t, err)
	require.Equal(t, []byte("bbb"), out)

	out, err = DecodeSingle(buf.Bytes(), "CCC")
	require.Error(t, err)
}

func TestEncodeSingle(t *testing.T) {
	buf := EncodeSingle([]byte("aaa"), "AAA")
	out, err := DecodeSingle(buf, "AAA")
	require.NoError(t, err)
	require.Equal(t, []byte("aaa"), out)
}
