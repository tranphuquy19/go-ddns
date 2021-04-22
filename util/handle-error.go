package util

import (
	"log"
	"strings"
)

func HandleError(err error, messages ...string) {
	if err != nil {
		messages = append(messages, "Error: "+err.Error())
		log.Fatal(strings.Join(messages, ". "))
	}
}
