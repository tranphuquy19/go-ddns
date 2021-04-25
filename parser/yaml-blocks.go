package parser

type Config struct {
	Providers []Provider `yaml:"providers" validate:"required,dive"`
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
	Name    string  `yaml:"name" validate:"required"`
	Type    string  `yaml:"type" validate:"required"`
	Source  Source  `yaml:"source" validate:"required,dive,required"`
	Trigger Trigger `yaml:"trigger" validate:"required,dive,required"`
	TTL     uint32  `yaml:"TTL" validate:"number,gt=0,max=2147483647"`
}

type Source struct {
	Value string `yaml:"value" validate:"required"`
	Type  string `yaml:"type" validate:"required"`
}

type Trigger struct {
	Value  string   `yaml:"value"`
	Type   string   `yaml:"type" validate:"required"`
	Values []string `yaml:"values"`
}
