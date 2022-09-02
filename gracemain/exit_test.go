package gracemain

import (
	"errors"
	"github.com/stretchr/testify/require"
	"os"
	"testing"
)

type testErrWithCode struct{}

func (testErrWithCode) Error() string {
	return "a"
}

func (testErrWithCode) ExitCode() int {
	return 8
}

func TestExit(t *testing.T) {
	var (
		savedErr  = errors.New("BBBBB")
		savedCode = -999
	)
	osExit = func(code int) {
		savedCode = code
	}
	OnExit(func(err *error) {
		savedErr = *err
	})
	defer func() {
		OnExit(DefaultOnExit)
		osExit = os.Exit
	}()
	var err error
	Exit(&err)
	require.Equal(t, 0, savedCode)
	require.Equal(t, nil, savedErr)
	err = testErrWithCode{}
	Exit(&err)
	require.Equal(t, 8, savedCode)
	require.Equal(t, err, savedErr)
}
