package gracek8s

import (
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"reflect"
)

// ExtractObjectMeta extract object 'metadata' field from any kubernetes resource object
func ExtractObjectMeta(v any) (meta v1.ObjectMeta) {
	val := reflect.ValueOf(v)
	for val.Kind() == reflect.Pointer {
		val = val.Elem()
	}
	if val.Kind() != reflect.Struct {
		return
	}
	val = val.FieldByName("ObjectMeta")
	if val.IsZero() {
		return
	}
	meta, _ = val.Interface().(v1.ObjectMeta)
	return
}
