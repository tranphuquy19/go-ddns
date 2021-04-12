package util

import (
	"log"
	"strings"
)

func HandleError(err error, messages ...string) {
	if err != nil {
		log.Fatal(strings.Join(messages, ". "))
	}
}
