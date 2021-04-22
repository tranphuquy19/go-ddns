package util

import (
	"fmt"
	"strings"
)

func ParseURL(baseUrl string, endpoint ...string) string {
	temp := append([]string{baseUrl}, endpoint...)
	url := strings.Join(temp, "/")
	return url
}

func ParseRecordURL(recordName string, domainStr string) string {
	url := ""
	if recordName != "@" {
		url = fmt.Sprintf("%s.%s", recordName, domainStr)
	} else {
		url = domainStr
	}
	return url
}
