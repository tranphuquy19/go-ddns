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

func DelRecordById(zoneId string, recordId string, record *parser.Record) gjson.Result {
	res, _ := netlifyClient.Del(fmt.Sprintf("dns_zone/%s/dns_records/%s", zoneId, recordId))
	return gjson.Parse(res)
}

func getIp(record *parser.Record, currentIP chan string) {
	ip := ""
	if record.Source.Type == "GET" || record.Source.Type == "POST" {
		client := InitClient(record.Source.Value, "", "")
		ip, _ = client.Get()
	}
	currentIP <- ip
}

func NetlifyUpdateRecord(domainName string, record *parser.Record, token string) {
	netlifyClient = InitNetlifyClient(token)
	url := util.ParseRecordURL(record.Name, domainName)
	currentIP := make(chan string)
	go getIp(record, currentIP)
	zones := GetDNSZones().Array()
	zoneId := ""
	for _, _zone := range zones {
		jZone := gjson.Parse(_zone.Raw)
		zoneName, _zoneId := jZone.Get("name").String(), jZone.Get("id").String()
		if zoneName == domainName {
			zoneId = _zoneId
			break
		}
	}
	crtIp := <- currentIP
	recordId, recordValue := "", ""
	if zoneId != "" && crtIp != "" {
		records := GetRecords(zoneId).Array()
		for _, _record := range records {
			jRecord := gjson.Parse(_record.Raw)
			_recordName, _recordId, _recordType, _recordValue := jRecord.Get("hostname").String(), jRecord.Get("id").String(), jRecord.Get("type").String(), jRecord.Get("value").String()
			fmt.Println(_recordName, _recordId, _recordType, _recordValue)
			if url == _recordName {
				recordId = _recordId
				break
			}
		}
	} else {
		util.HandleError(nil, fmt.Sprintf("Site not found: %s", url))
		return
	}

	if crtIp != "" {
		if crtIp != recordValue {
			_ = DelRecordById(zoneId, recordId, record)
		}
	} else {
		util.HandleError(nil, "Cannot get your public IP")
		return
	}
}
