package gracek8s

import (
	"github.com/stretchr/testify/require"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"testing"
)

func TestExtractObjectMeta(t *testing.T) {
	metadata := ExtractObjectMeta(&corev1.Service{ObjectMeta: metav1.ObjectMeta{
		Name:      "aaa",
		Namespace: "bbb",
	}})
	require.Equal(t, "aaa", metadata.Name)
	require.Equal(t, "bbb", metadata.Namespace)
	metadata = ExtractObjectMeta(corev1.ConfigMap{ObjectMeta: metav1.ObjectMeta{
		Name:      "aaa",
		Namespace: "bbb",
	}})
	require.Equal(t, "aaa", metadata.Name)
	require.Equal(t, "bbb", metadata.Namespace)
}
