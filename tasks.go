package main

import (
	"fmt"
	client "go-ddns/clients"
)

func getIP(baseUrl string, method string) {
	// test http-client - Get current public IP
	client := client.InitClient(baseUrl, "", "")
	res, _ := client.Get()
	fmt.Println("Your IP: ", res)
}
