package parser

import (
	"go-ddns/util"
	"io/ioutil"

	"gopkg.in/yaml.v2"
)

func YAMLParser(filePath string) Config {
	config := Config{}

	yamlFile, e1 := ioutil.ReadFile(filePath)

	util.HandleError(e1, "An error occurred while reading the YAML file")

	e2 := yaml.Unmarshal(yamlFile, &config)
	util.HandleError(e2, "An error occurred while passing the YAML file")

	YAMLValidator(config)

	return config
}
