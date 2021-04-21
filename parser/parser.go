package parser

import (
	"fmt"
	"go-ddns/util"
	"io/ioutil"

	"github.com/pelletier/go-toml"
	"gopkg.in/yaml.v2"
)

func ConfigYAMLParser(filePath string) Config {
	config := Config{}

	yamlFile, e1 := ioutil.ReadFile(filePath)

	util.HandleError(e1, "An error occurred while reading the YAML file")

	e2 := yaml.Unmarshal(yamlFile, &config)
	util.HandleError(e2, "An error occurred while passing the YAML file")

	YAMLValidator(config)

	return config
}

func TOMLGetProfile(filePath string, profile string) string {
	tomlContent, err := ioutil.ReadFile(filePath)
	util.HandleError(err, "An error occurred while reading the Credentials file")

	tomlFile, _ := toml.Load(string(tomlContent))

	return tomlFile.Get(fmt.Sprintf("%s.TOKEN", profile)).(string)
}
