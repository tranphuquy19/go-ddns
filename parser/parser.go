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

	yamlFile, err := ioutil.ReadFile(filePath)
	util.HandleError(err, "An error occurred while reading the YAML file", "File: "+filePath)

	err = yaml.Unmarshal(yamlFile, &config)
	util.HandleError(err, "An error occurred while parsing the YAML file", "File: "+filePath)
	YAMLValidator(config)

	return config
}

func TOMLGetProfile(filePath string, profile string) string {
	tomlContent, err := ioutil.ReadFile(filePath)
	util.HandleError(err, "An error occurred while reading the Credentials file")

	tomlFile, _ := toml.Load(string(tomlContent))

	return tomlFile.Get(fmt.Sprintf("%s.TOKEN", profile)).(string)
}
