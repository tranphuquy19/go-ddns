package client

import "github.com/tidwall/gjson"

func InitNetlifyClient(token string) HttpClient {
	netlifyClient := InitClient("https://api.netlify.com/api/v1", token, "Bearer")
	return *netlifyClient
}

func GetDNSZones(client HttpClient) gjson.Result {
	res, _ := client.Get()
	return gjson.Get(res, "#.id")
}
