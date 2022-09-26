package gracek8s

import (
	"context"
	"errors"
	"fmt"
	"github.com/guoyk93/grace/gracex509"
	corev1 "k8s.io/api/core/v1"
	kerrors "k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
)

type APIGetCreate[T any] interface {
	Get(ctx context.Context, name string, opts metav1.GetOptions) (*T, error)
	Create(ctx context.Context, obj *T, opts metav1.CreateOptions) (*T, error)
}

func Ensure[T any](ctx context.Context, api APIGetCreate[T], obj *T) (out *T, err error) {
	metadata := ExtractObjectMeta(obj)

	if metadata.Name == "" {
		err = errors.New("gracek8s.Ensure: missing metadata.name")
		return
	}

	if out, err = api.Get(ctx, metadata.Name, metav1.GetOptions{}); err != nil {
		if kerrors.IsNotFound(err) {
			out, err = api.Create(ctx, obj, metav1.CreateOptions{})
		}
	}

	return
}

type EnsureCertificateOptions struct {
	Name      string
	Namespace string
	gracex509.GenerationOptions
}

type EnsureCertificateResult struct {
	CrtPEM []byte
	KeyPEM []byte
}

func EnsureCertificate(ctx context.Context, client *kubernetes.Clientset, opts EnsureCertificateOptions) (res EnsureCertificateResult, err error) {
	var secret *corev1.Secret
	if secret, err = client.CoreV1().Secrets(opts.Namespace).Get(ctx, opts.Name, metav1.GetOptions{}); err != nil {
		if kerrors.IsNotFound(err) {
			err = nil

			var xRes gracex509.GenerationResult

			if xRes, err = gracex509.Generate(opts.GenerationOptions); err != nil {
				return
			}

			res.CrtPEM = xRes.CrtPEM
			res.KeyPEM = xRes.KeyPEM

			if _, err = client.CoreV1().Secrets(opts.Namespace).Create(ctx, &corev1.Secret{
				ObjectMeta: metav1.ObjectMeta{
					Name: opts.Name,
				},
				Type: corev1.SecretTypeTLS,
				StringData: map[string]string{
					corev1.TLSCertKey:       string(res.CrtPEM),
					corev1.TLSPrivateKeyKey: string(res.KeyPEM),
				},
			}, metav1.CreateOptions{}); err != nil {
				return
			}
			return
		} else {
			return
		}
	} else {
		res.CrtPEM, res.KeyPEM = secret.Data[corev1.TLSCertKey], secret.Data[corev1.TLSPrivateKeyKey]
		if len(res.CrtPEM) == 0 {
			err = fmt.Errorf("missing key: %s", corev1.TLSCertKey)
			return
		}
		if len(res.KeyPEM) == 0 {
			err = fmt.Errorf("missing key: %s", corev1.TLSPrivateKeyKey)
			return
		}
	}
	return
}
