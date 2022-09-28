package gracek8s

import (
	"context"
	"github.com/guoyk93/grace/gracex509"
	"github.com/stretchr/testify/require"
	corev1 "k8s.io/api/core/v1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"testing"
)

func TestGetOrCreate(t *testing.T) {
	if ShouldSkip() {
		return
	}

	client, err := DefaultClient()
	require.NoError(t, err)

	cm, err := GetOrCreate[corev1.ConfigMap](
		context.Background(),
		client.CoreV1().ConfigMaps("default"),
		&corev1.ConfigMap{
			ObjectMeta: v1.ObjectMeta{
				Name: "test-get-or-create",
				Annotations: map[string]string{
					"aaa": "bbb",
				},
			},
		},
	)
	require.NoError(t, err)
	require.Equal(t, "bbb", cm.Annotations["aaa"])

	cm, err = GetOrCreate[corev1.ConfigMap](
		context.Background(),
		client.CoreV1().ConfigMaps("default"),
		&corev1.ConfigMap{
			ObjectMeta: v1.ObjectMeta{
				Name: "test-get-or-create",
				Annotations: map[string]string{
					"aaa": "ccc",
				},
			},
		},
	)
	require.NoError(t, err)
	require.Equal(t, "bbb", cm.Annotations["aaa"])
}

func TestGetOrCreateTLSSecret(t *testing.T) {
	if ShouldSkip() {
		return
	}

	client, err := DefaultClient()
	require.NoError(t, err)

	_, sec, err := GetOrCreateTLSSecret(
		context.Background(),
		client.CoreV1().Secrets("default"),
		"test-get-or-create-tls",
		gracex509.GenerateOptions{
			IsCA:  true,
			Names: []string{"test-get-or-create-tls"},
		},
	)
	require.NoError(t, err)
	require.False(t, sec.IsZero())
}
