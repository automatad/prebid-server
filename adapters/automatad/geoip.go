package automatad

import (
	"fmt"
	"net"

	"github.com/oschwald/maxminddb-golang"
)

var DB *maxminddb.Reader
var err error

type CountryRecord struct {
	Country struct {
		ISOCode string `maxminddb:"iso_code"`
	} `maxminddb:"country"`
}

func init() {
	DB, err = maxminddb.Open("adapters/automatad/GeoLite2-Country.mmdb")
}

func CountryFromIP(ip_string string) (string, error) {

	if DB == nil {
		return "", fmt.Errorf("recieved following error while loading maxmind country db: %s", err.Error())
	}

	ip := net.ParseIP(ip_string)

	record := CountryRecord{}

	err := DB.Lookup(ip, &record)
	if err != nil {
		return "", err
	}
	return record.Country.ISOCode, nil
}
