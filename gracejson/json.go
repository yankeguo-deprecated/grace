package gracejson

import (
	"encoding/json"
	"os"
)

func UnmarshalJSON[T any](buf []byte) (out T, err error) {
	err = json.Unmarshal(buf, &out)
	return
}

func UnmarshalJSONFile[T any](filename string) (T, error) {
	if buf, err := os.ReadFile(filename); err != nil {
		var v T
		return v, err
	} else {
		return UnmarshalJSON[T](buf)
	}
}

func MarshalPretty(v interface{}) ([]byte, error) {
	return json.MarshalIndent(v, "", "  ")
}
