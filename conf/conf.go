package conf

import (
	"encoding/json"
	"encoding/xml"
	"github.com/BurntSushi/toml"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"log"
	"strings"
)

func InitConf(path string, cfg interface{}) interface{} {
	byt, err := ioutil.ReadFile(path)
	if err != nil {
		panic(err)
	}
	pathArr := strings.Split(path, ".")
	suffix := pathArr[len(pathArr)-1]
	switch suffix {
	case "yml":
		err = yaml.Unmarshal(byt, cfg)
	case "json":
		err = json.Unmarshal(byt, cfg)
	case "toml":
		_, err = toml.DecodeFile(path, cfg)
	case "xml":
		err = xml.Unmarshal(byt, cfg)
	default:
		log.Panicf("config file format (%s) do not support", suffix)
	}
	if err != nil {
		panic(err)
	}
	return cfg
}
