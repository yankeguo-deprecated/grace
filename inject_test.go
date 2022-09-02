package grace

import (
	"context"
	"testing"

	"github.com/stretchr/testify/require"
)

type testInjectType struct {
	A string
}

func (ti testInjectType) B() string {
	return ti.A
}

type testInjectInterface interface {
	B() string
}

func TestInject(t *testing.T) {
	vt := InjectKey(&testInjectType{})
	require.Equal(t, injectKey("github.com/guoyk93/grace::testInjectType"), vt)

	v := testInjectType{A: "hello"}
	ctx := context.Background()
	ctx = Inject(ctx, v)
	v2, ok := Extract[testInjectType](ctx)
	require.True(t, ok)
	require.Equal(t, "hello", v2.A)
	ctx = Inject[testInjectInterface](ctx, v)
	v3, ok := Extract[testInjectInterface](ctx)
	require.True(t, ok)
	require.Equal(t, "hello", v3.B())
}
