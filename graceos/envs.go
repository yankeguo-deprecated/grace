package graceos

import (
	"errors"
	"os"
	"reflect"
	"strconv"
	"strings"
	"time"

	"github.com/guoyk93/grace"
)

var (
	EnvPrefix = ""
)

func EnvVal[T any](out *T, key string, required bool) grace.TaskFunc {
	prefix := EnvPrefix
	return func() error {
		val := strings.TrimSpace(os.Getenv(prefix + key))
		if val == "" {
			if required {
				return errors.New("missing environment variable $" + EnvPrefix + key)
			}
			return nil
		} else {
			return DecodeEnvVal(out, val)
		}
	}
}

func EnvSlice[T any](out *[]T, key string, required bool) grace.TaskFunc {
	prefix := EnvPrefix
	return func() error {
		val := strings.TrimSpace(os.Getenv(prefix + key))
		if val == "" {
			if required {
				return errors.New("missing environment variable $" + EnvPrefix + key)
			}
			return nil
		} else {
			return DecodeEnvSlice(out, val)
		}
	}
}

func EnvMap[T any](out *map[string]T, key string, required bool) grace.TaskFunc {
	prefix := EnvPrefix
	return func() error {
		val := strings.TrimSpace(os.Getenv(prefix + key))
		if val == "" {
			if required {
				return errors.New("missing environment variable $" + EnvPrefix + key)
			}
			return nil
		} else {
			return DecodeEnvMap(out, val)
		}
	}
}

func DecodeEnvVal[T any](out *T, s string) (err error) {
	var o interface{} = out
	switch o.(type) {
	case *string:
		if strings.HasPrefix(s, "\"") && strings.HasSuffix(s, "\"") {
			var v string
			if v, err = strconv.Unquote(s); err != nil {
				return
			}
			*(o).(*string) = v
		} else {
			*(o).(*string) = s
		}
	case *bool:
		var v bool
		if v, err = strconv.ParseBool(s); err != nil {
			return
		}
		*(o).(*bool) = v
	case *int:
		var v int
		if v, err = strconv.Atoi(s); err != nil {
			return
		}
		*(o).(*int) = v
	case *int64:
		var v int64
		if v, err = strconv.ParseInt(s, 10, 64); err != nil {
			return
		}
		*(o).(*int64) = v
	case *uint64:
		var v uint64
		if v, err = strconv.ParseUint(s, 10, 64); err != nil {
			return
		}
		*(o).(*uint64) = v
	case *float64:
		var v float64
		if v, err = strconv.ParseFloat(s, 64); err != nil {
			return
		}
		*(o).(*float64) = v
	case *time.Duration:
		var v time.Duration
		if v, err = time.ParseDuration(s); err != nil {
			return
		}
		*(o).(*time.Duration) = v
	default:
		err = errors.New("DecodeEnvVal: unsupported type: " + reflect.TypeOf(o).Elem().Name())
	}
	return
}

func DecodeEnvSlice[T any](out *[]T, s string) (err error) {
	data := []T{}
	splits := strings.Split(s, ",")
	for _, split := range splits {
		var v T
		if err = DecodeEnvVal(&v, split); err != nil {
			return
		}
		data = append(data, v)
	}
	*out = data
	return
}

func DecodeEnvMap[T any](out *map[string]T, s string) (err error) {
	data := map[string]T{}
	splits := strings.Split(s, ";")
	for _, split := range splits {
		kvs := strings.SplitN(split, ":", 2)
		if len(kvs) != 2 {
			continue
		}
		var k string
		if err = DecodeEnvVal(&k, kvs[0]); err != nil {
			return
		}
		var v T
		if err = DecodeEnvVal(&v, kvs[1]); err != nil {
			return
		}
		data[k] = v
	}
	*out = data
	return
}
