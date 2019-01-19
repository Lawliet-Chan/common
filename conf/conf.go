package conf

import (
	"gopkg.in/yaml.v2"
	"io/ioutil"
)

func InitConf(path string, cfg interface{}) interface{} {
	byt, err := ioutil.ReadFile(path)
	if err != nil {
		panic(err)
	}
	err = yaml.Unmarshal(byt, cfg)
	if err != nil {
		panic(err)
	}
	return cfg
}
