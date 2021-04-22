package client

import (
	"fmt"
	parser "go-ddns/parser"
	"go-ddns/util"

	"github.com/tidwall/gjson"
)

var client HttpClient

func InitNetlifyClient(token string) HttpClient {
	netlifyClient := InitClient("https://api.netlify.com/api/v1", token, "Bearer")
	return *netlifyClient
}

func GetDNSZones() gjson.Result {
	res, _ := client.Get("dns_zones")
	return gjson.Parse(res)
}

func GetRecords(zoneId string) gjson.Result {
	res, _ := client.Get(fmt.Sprintf("dns_zones/%s/dns_records", zoneId))
	fmt.Println(zoneId, "=================")
	return gjson.Parse(res)
}

func GetRecordById(zoneId string, recordId string, record *parser.Record) gjson.Result {
	res, _ := client.Get(fmt.Sprintf("dns_zone/%s/dns_records/%s", zoneId, recordId))
	return gjson.Parse(res)
}

func NetlifyUpdateRecord(domainName string, record *parser.Record, token string) {
	client = InitNetlifyClient(token)
	url := util.ParseRecordURL(record.Name, domainName)
	fmt.Println(url)
	GetDNSZones().ForEach(func(zoneKey, zone gjson.Result) bool {
		domain := gjson.Get(zone.String(), "name").String()
		zoneId := gjson.Get(zone.String(), "id").String()
		fmt.Println("Domain:", domain, ",", "zoneId:", zoneId)
		if domain == domainName {
			GetRecords(zoneId).ForEach(func(recordKey, jRecord gjson.Result) bool {
				recordId := jRecord.Get("id").String()
				fmt.Println(recordKey, recordId)
				return true
			})
			return true // keep iterating
		} else {
			return false // stop iterating
		}
	})
}
