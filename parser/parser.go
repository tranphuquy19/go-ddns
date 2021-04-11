package parser

import (
	"fmt"
	"go-ddns/util"
	"io/ioutil"

	"gopkg.in/yaml.v2"
)

func YAMLParser(filePath string) {
	config := Config{}

	yamlFile, e1 := ioutil.ReadFile(filePath)

	util.HandleError("An error occurred while reading the YAML file", e1)

	e2 := yaml.Unmarshal(yamlFile, &config)

	util.HandleError("An error occurred while passing the YAML file", e2)

	fmt.Println(config.Providers[0].Domains[0].Records[0].Value)

	fmt.Println(filePath)
}
