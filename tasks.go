package main

import (
	client "go-ddns/clients"
)

func getIP(baseUrl string, method string) string {
	// test http-client - Get current public IP
	client := client.InitClient(baseUrl, "", "")
	currentIP, _ := client.Get()
	return currentIP
}
