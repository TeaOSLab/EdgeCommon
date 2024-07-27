package configutils

import (
	"os"

	"gopkg.in/yaml.v3"
)

func UnmarshalYamlFile(file string, ptr interface{}) error {
	data, err := os.ReadFile(file)
	if err != nil {
		return err
	}
	return yaml.Unmarshal(data, ptr)
}
