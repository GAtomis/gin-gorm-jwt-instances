package common

import (
	"github.com/BurntSushi/toml"
)

var Config map[string]interface{} = make(map[string]interface{}, 0)

func LoadConfig(path string) error {
	_, err := toml.DecodeFile(path, &Config)
	return err
}

func GetMustStringValue(key string, name string) string {
	return Config[key].(map[string]interface{})[name].(string)

}

func GetMustIntValue(key string, name string) int {
	return int(Config[key].(map[string]interface{})[name].(int64))
}
