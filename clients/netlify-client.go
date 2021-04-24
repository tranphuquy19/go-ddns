package client

import (
	"fmt"
	parser "go-ddns/parser"
	"go-ddns/util"

	"github.com/tidwall/gjson"
)

var netlifyClient HttpClient

type NetlifyRecord struct {
	Hostname string
	Type     string
	TTL      int64
	Id       string
	Value    string
}

func InitNetlifyClient(token string) HttpClient {
	netlifyClient := InitClient("https://api.netlify.com/api/v1", token, "Bearer")
	return *netlifyClient
}

func CreateRecord(zoneId string, values []byte) gjson.Result {
	res, _ := netlifyClient.Post(values, fmt.Sprintf("dns_zones/%s/dns_records", zoneId))
	return gjson.Parse(res)
}

func GetDNSZones() gjson.Result {
	res, _ := netlifyClient.Get(nil, "dns_zones")
	return gjson.Parse(res)
}

func GetRecords(zoneId string) gjson.Result {
	res, _ := netlifyClient.Get(nil, fmt.Sprintf("dns_zones/%s/dns_records", zoneId))
	return gjson.Parse(res)
}

func GetRecordById(zoneId string, recordId string, record *parser.Record) gjson.Result {
	res, _ := netlifyClient.Get(nil, fmt.Sprintf("dns_zone/%s/dns_records/%s", zoneId, recordId))
	return gjson.Parse(res)
}

func DelRecordById(zoneId string, recordId string, record *parser.Record) gjson.Result {
	res, _ := netlifyClient.Del(nil, fmt.Sprintf("dns_zone/%s/dns_records/%s", zoneId, recordId))
	return gjson.Parse(res)
}

func getIp(record *parser.Record, currentIP chan string) {
	ip := ""
	if record.Source.Type == "GET" || record.Source.Type == "POST" {
		client := InitClient(record.Source.Value, "", "")
		ip, _ = client.Get(nil)
	}
	currentIP <- ip
}

func getRecord(zoneId string, url string) NetlifyRecord {
	var gotRecord = NetlifyRecord{}
	if zoneId != "" {
		records := GetRecords(zoneId).Array()
		for _, _record := range records {
			jRecord := gjson.Parse(_record.Raw)
			_recordName, _recordId, _recordType, _recordValue, _recordTTL := jRecord.Get("hostname").String(), jRecord.Get("id").String(), jRecord.Get("type").String(), jRecord.Get("value").String(), jRecord.Get("ttl").Int()
			fmt.Println(_recordName, _recordId, _recordType, _recordValue)
			if url == _recordName {
				gotRecord = NetlifyRecord{
					Hostname: _recordName,
					Type:     _recordType,
					TTL:      _recordTTL,
					Id:       _recordId,
					Value:    _recordValue,
				}
				break
			}
		}
		return gotRecord
	} else {
		util.HandleError(nil, fmt.Sprintf("Site not found: %s", url))
		return gotRecord
	}
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
	var gotRecord = getRecord(zoneId, url)

	crtIp := <-currentIP
	if crtIp != "" {
		if crtIp != gotRecord.Value {
			DelRecordById(zoneId, gotRecord.Id, record)
			gotRecord.Value = crtIp
			jsonNewRecord := fmt.Sprintf(`{ "hostname": "%s", "type": "%s", "ttl": %s, "value": "%s" }`, gotRecord.Hostname, gotRecord.Type, gotRecord.Type, gotRecord.Value)
			values := []byte(jsonNewRecord)
			CreateRecord(zoneId, values)
		}
	} else {
		util.HandleError(nil, "Cannot get your public IP")
		return
	}

}
