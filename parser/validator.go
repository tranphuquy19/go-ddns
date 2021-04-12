package parser

import (
	"fmt"
	"go-ddns/util"
	"strings"

	"github.com/go-playground/validator/v10"
)

var validate *validator.Validate

func YAMLValidator(config Config) {
	validate = validator.New()

	validate.RegisterStructValidation(recordValidator, Config{})

	err := validate.Struct(config)

	if err != nil {
		if _, ok := err.(*validator.InvalidValidationError); ok {
			util.HandleError(err, "Invalid YAML file")
			return
		}

		for _, err := range err.(validator.ValidationErrors) {
			errValue := fmt.Sprintf("%v", err.Value())
			util.HandleError(err, "Invalid "+err.Tag()+": "+errValue)
		}

		// from here you can create your own error messages in whatever language you wish
		return
	}
}

func recordValidator(sl validator.StructLevel) {
	recordTypes := []string{"ipv4", "ipv6", "get", "post", "nslookup", "txt"}
	config := sl.Current().Interface().(Config)

	for _, provider := range config.Providers {
		for _, domain := range provider.Domains {
			for _, record := range domain.Records {
				lwRecordStr := strings.ToLower(record.Type)

				// check record type
				if !util.Contains(recordTypes, lwRecordStr) {
					sl.ReportError(record.Type, "type", "Type", "record_type", "")
				}

				// check record value
				var valueErr error
				switch lwRecordStr {
				case "ipv4":
					valueErr = validate.Var(record.Value, "ip4_addr")
				case "ipv6":
					valueErr = validate.Var(record.Value, "ip6_addr")
				case "txt":
					valueErr = validate.Var(record.Value, "ascii")
				case "get", "post", "nslookup":
					valueErr = validate.Var(record.Value, "uri")
				}

				if valueErr != nil {
					sl.ReportError(record.Value, "type", "Value", "record_value", "")
				}
			}
		}
	}
}
