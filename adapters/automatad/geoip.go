package automatad

import (
	"net"

	"github.com/oschwald/maxminddb-golang"
)

func InitDB() (*maxminddb.Reader, error) {
	db, err := maxminddb.Open("adapters/automatad/GeoLite2-Country.mmdb")
	if err != nil {
		return nil, err
	}
	return db, nil
}

var record struct {
	Country struct {
		ISOCode string `maxminddb:"iso_code"`
	} `maxminddb:"country"`
}

func CountryFromIP(db *maxminddb.Reader, ip_string string) (string, error) {
	//db, err := maxminddb.Open("test-data/test-data/GeoIP2-City-Test.mmdb")
	//if err != nil {
	//	log.Fatal(err)
	//}
	//defer db.Close()

	ip := net.ParseIP(ip_string)

	//var recordstruct {
	//	Country struct {
	//		ISOCode string `maxminddb:"iso_code"`
	//	} `maxminddb:"country"`
	//} // Or any appropriate struct

	err := db.Lookup(ip, &record)
	if err != nil {
		return "", err
	}
	return record.Country.ISOCode, nil
}
