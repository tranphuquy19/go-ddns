package client

import (
	"fmt"
	parser "go-ddns/parser"
	"go-ddns/util"

	"github.com/tidwall/gjson"
)

var netlifyClient HttpClient

func InitNetlifyClient(token string) HttpClient {
	netlifyClient := InitClient("https://api.netlify.com/api/v1", token, "Bearer")
	return *netlifyClient
}

func GetDNSZones() gjson.Result {
	res, _ := netlifyClient.Get("dns_zones")
	return gjson.Parse(res)
}

func GetRecords(zoneId string) gjson.Result {
	res, _ := netlifyClient.Get(fmt.Sprintf("dns_zones/%s/dns_records", zoneId))
	return gjson.Parse(res)
}

func GetRecordById(zoneId string, recordId string, record *parser.Record) gjson.Result {
	res, _ := netlifyClient.Get(fmt.Sprintf("dns_zone/%s/dns_records/%s", zoneId, recordId))
	return gjson.Parse(res)
}

func NetlifyUpdateRecord(domainName string, record *parser.Record, token string) {
	netlifyClient = InitNetlifyClient(token)
	url := util.ParseRecordURL(record.Name, domainName)
	fmt.Println(url)
	zones := GetDNSZones().Array()
	for _, _zone := range zones {
		jZone := gjson.Parse(_zone.Raw)
		zoneName, zoneId := jZone.Get("name").String(), jZone.Get("id").String()
		if zoneName == domainName {
			records := GetRecords(zoneId).Array()
			for _, _record := range records {
				jRecord := gjson.Parse(_record.Raw)
				recordName, recordId, recordType, recordValue := jRecord.Get("hostname").String(), jRecord.Get("id").String(), jRecord.Get("type").String(), jRecord.Get("value").String()
				fmt.Println(recordName, recordId, recordType, recordValue)
			}
			break
		}
	}
}
