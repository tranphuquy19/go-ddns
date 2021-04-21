package client

import (
	"fmt"
	parser "go-ddns/parser"

	"github.com/tidwall/gjson"
)

func InitNetlifyClient(token string) HttpClient {
	netlifyClient := InitClient("https://api.netlify.com/api/v1", token, "Bearer")
	return *netlifyClient
}

func GetDNSZones(client HttpClient) gjson.Result {
	res, _ := client.Get("dns_zones")
	return gjson.Parse(res)
}

func NetlifyUpdateRecord(domainName string, record *parser.Record, token string) {
	client := InitNetlifyClient(token)
	GetDNSZones(client).ForEach(func(key, zone gjson.Result) bool {
		domain := gjson.Get(zone.String(), "name")
		id := gjson.Get(zone.String(), "id")
		fmt.Println("Domain:", domain, ",", "id:", id)
		return true // keep iterating
	})
}
