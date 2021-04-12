package parser

type Config struct {
	Providers []Provider `yaml:"provider"`
}

type Provider struct {
	Name    string   `yaml:"name"`
	Profile string   `yaml:"profile"`
	Domains []Domain `yaml:"domains"`
}

type Domain struct {
	Name    string   `yaml:"name"`
	Records []Record `yaml:"records"`
}

type Record struct {
	Name  	string `yaml:"name"`
	Value 	string `yaml:"value"`
	Type 	string `yaml:"type"`
}
