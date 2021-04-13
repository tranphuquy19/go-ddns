package parser

import (
	"fmt"
	client "go-ddns/clients"
	"go-ddns/util"
	"io/ioutil"

	"gopkg.in/yaml.v2"
)

func YAMLParser(filePath string) {
	config := Config{}

	yamlFile, e1 := ioutil.ReadFile(filePath)

	util.HandleError(e1, "An error occurred while reading the YAML file")

	e2 := yaml.Unmarshal(yamlFile, &config)
	util.HandleError(e2, "An error occurred while passing the YAML file")

	YAMLValidator(config)

	// test http-client - Get current public IP
	baseUrl := "http://api.ipify.org"
	client := client.InitClient(baseUrl, "", "")
	res, _ := client.Get()
	fmt.Println("Your IP: ", res)
}
