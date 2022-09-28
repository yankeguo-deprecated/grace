package gracek8s

import (
	"context"
	"errors"
	"fmt"
	"github.com/guoyk93/grace/gracex509"
	corev1 "k8s.io/api/core/v1"
	kerrors "k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

type APIGetCreate[T any] interface {
	Get(ctx context.Context, name string, opts metav1.GetOptions) (*T, error)
	Create(ctx context.Context, obj *T, opts metav1.CreateOptions) (*T, error)
}

// GetOrCreate get or create a kubernetes resource
func GetOrCreate[T any](ctx context.Context, api APIGetCreate[T], obj *T) (out *T, err error) {
	metadata := ExtractMetadata(obj)

	if metadata.Name == "" {
		err = errors.New("gracek8s.GetOrCreate: missing metadata.name")
		return
	}

	if out, err = api.Get(ctx, metadata.Name, metav1.GetOptions{}); err != nil {
		if kerrors.IsNotFound(err) {
			out, err = api.Create(ctx, obj, metav1.CreateOptions{})
		}
	}

	return
}

// GetOrCreateTLSSecret get or create a secret with type tls, using gracex509.Generate
func GetOrCreateTLSSecret(
	ctx context.Context,
	api APIGetCreate[corev1.Secret],
	name string,
	opts gracex509.GenerateOptions,
) (res gracex509.PEMPair, err error) {
	var secret *corev1.Secret
	if secret, err = api.Get(ctx, name, metav1.GetOptions{}); err != nil {
		if kerrors.IsNotFound(err) {
			if res, err = gracex509.Generate(opts); err != nil {
				return
			}

			if _, err = api.Create(ctx, &corev1.Secret{
				ObjectMeta: metav1.ObjectMeta{
					Name: name,
				},
				Type: corev1.SecretTypeTLS,
				StringData: map[string]string{
					corev1.TLSCertKey:       string(res.Crt),
					corev1.TLSPrivateKeyKey: string(res.Key),
				},
			}, metav1.CreateOptions{}); err != nil {
				return
			}
		}
	} else {
		res.Crt, res.Key = secret.Data[corev1.TLSCertKey], secret.Data[corev1.TLSPrivateKeyKey]

		if res.IsZero() {
			err = fmt.Errorf("missing key: %s or %s", corev1.TLSCertKey, corev1.TLSPrivateKeyKey)
			return
		}
	}
	return
}
