package utils

import (
	"avitocalls/internal/pkg/forms"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

var (
	err      error
	geo      forms.GeoIP
	response *http.Response
	body     []byte
)

func GetByIP(address string) (forms.GeoIP, error) {
	response, err = http.Get("http://free.ipwhois.io/json/" + address)
	if err != nil {
		fmt.Println(err)
	}
	defer response.Body.Close()
	body, err = ioutil.ReadAll(response.Body)
	if err != nil {
		return geo, err
	}
	err = json.Unmarshal(body, &geo)
	return geo, err

	//{"ip":"217.175.11.137","success":true,
	//"type":"IPv4","continent":"Europe","continent_code":"EU",
	//"country":"Russia","country_code":"RU","country_flag":"https:\/\/cdn.ipwhois.io\/flags\/ru.svg",
	//"country_capital":"Moscow","country_phone":"+7","country_neighbours":"GE,CN,BY,UA,KZ,LV,PL,EE,LT,FI,MN,NO,AZ,KP",
	//"region":"Krasnodar Krai","city":"Anapa","latitude":"44.8857008","longitude":"37.3199191",
	//"asn":"AS39047","org":"KerchNET-PLUS","isp":"Multiservice Networks Ltd.",
	//"timezone":"Europe\/Moscow","timezone_name":"Moscow Standard Time","timezone_dstOffset":"0",
	//"timezone_gmtOffset":"10800","timezone_gmt":"GMT +3:00","currency":"Russian Ruble","currency_code":"RUB",
	//"currency_symbol":"\u20bd","currency_rates":"74.1836","currency_plural":"Russian rubles","completed_requests":1}
}
