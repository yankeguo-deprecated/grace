package gracek8s

import (
	"encoding/json"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"reflect"
)

type MetadataObject struct {
	v1.ObjectMeta `json:"metadata"`
}

// ExtractMetadata extract 'metadata' field from any kubernetes resource object
func ExtractMetadata(v any) v1.ObjectMeta {
	val := reflect.ValueOf(v)
	for val.Kind() == reflect.Pointer {
		val = val.Elem()
	}
	if val.Kind() != reflect.Struct {
		goto modeJSON
	}
	if val = val.FieldByName("ObjectMeta"); val.IsZero() {
		goto modeJSON
	}
	if meta, ok := val.Interface().(v1.ObjectMeta); ok {
		return meta
	}
modeJSON:
	var obj MetadataObject
	if buf, err := json.Marshal(v); err == nil {
		_ = json.Unmarshal(buf, &obj)
	}
	return obj.ObjectMeta
}
