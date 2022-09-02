package graceconf

import (
	"flag"

	"github.com/creasty/defaults"
	"github.com/go-playground/validator/v10"
	"github.com/guoyk93/grace/graceyaml"
)

func LoadYAMLFlagConf[T any]() (out T, err error) {
	var conf string

	flag.StringVar(&conf, "conf", "config.yaml", "config file")
	flag.Parse()

	if out, err = graceyaml.UnmarshalYAMLFile[T](conf); err != nil {
		return
	}

	if err = defaults.Set(&out); err != nil {
		return
	}

	if err = validator.New().Struct(&out); err != nil {
		return
	}

	return
}
