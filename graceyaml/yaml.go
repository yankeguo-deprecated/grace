package graceyaml

import (
	"gopkg.in/yaml.v3"
	"io/ioutil"
)

func UnmarshalYAML[T any](buf []byte) (out T, err error) {
	err = yaml.Unmarshal(buf, &out)
	return
}

func UnmarshalYAMLFile[T any](filename string) (T, error) {
	if buf, err := ioutil.ReadFile(filename); err != nil {
		var v T
		return v, err
	} else {
		return UnmarshalYAML[T](buf)
	}
}
