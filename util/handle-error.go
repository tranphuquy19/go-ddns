package util

import "log"

func HandleError(message string, err error) {
	if err != nil {
		log.Fatal(message, err)
	}
}
