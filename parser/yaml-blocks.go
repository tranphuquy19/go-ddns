package parser

type Config struct {
	Providers []Provider `yaml:"provider" validate:"required,dive"`
}

type Provider struct {
	Name    string   `yaml:"name" validate:"required"`
	Profile string   `yaml:"profile" validate:"required"`
	Domains []Domain `yaml:"domains" validate:"required,dive,required"`
}

type Domain struct {
	Name    string   `yaml:"name" validate:"required"`
	Records []Record `yaml:"records" validate:"required,dive,required"`
}

type Record struct {
	Name  string `yaml:"name" validate:"required"`
	Value string `yaml:"value" validate:"required"`
	Type  string `yaml:"type" validate:"required"`
}
