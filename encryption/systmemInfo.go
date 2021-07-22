package helper

import (
	"fmt"
	"io/ioutil"
)

type systemInfo struct {
	DeviceName      string
	PlatformName    string
	PlatformVersion string
	OS              string
	OSVersion       string
	Country         string
	CountrISOCode   string
}

// https://www.maxmind.com/en/accounts/542317/geoip/downloads
func GetSystemInfo(useragent string, clientIP string) []string {

	/*db, err := geoip2.Open("./encryption/GeoLite2City.mmdb")

	if err != nil {
		fmt.Println(err)
	}
	defer db.Close()

	ip := net.ParseIP(clientIP)
	record, err := db.City(ip)

	if err != nil {
		fmt.Println(err)
	}

	ua := ua.Parse(useragent)

	var deviceSort string

	if ua.Mobile {
		deviceSort = "Mobile"
	}
	if ua.Tablet {
		deviceSort = "Tablet"
	}
	if ua.Desktop {
		deviceSort = "Desktop"
	}
	if ua.Bot {
		deviceSort = "Bot"
	}

	info := systemInfo{
		DeviceName:      deviceSort,
		PlatformName:    ua.Name,
		PlatformVersion: ua.Version,
		OS:              ua.OS,
		OSVersion:       ua.OSVersion,
		Country:         record.Country.Names["en"],
		CountrISOCode:   record.Country.IsoCode,
	}*/

	/*fmt.Println(ua)
	fmt.Println(ua.Name)
	fmt.Println(ua.Version)
	fmt.Println(ua.OS)
	fmt.Println(ua.OSVersion)
	fmt.Printf("Russian country name: %v\n", record.Country.Names["en"])
	fmt.Printf("ISO country code: %v\n", record.Country.IsoCode)
	fmt.Printf("Time zone: %v\n", record.Location.TimeZone)
	fmt.Printf("Coordinates: %v, %v\n", record.Location.Latitude, record.Location.Longitude)*/

	//return info

	files, _ := ioutil.ReadDir(".")

	var fileNames []string

	for _, file := range files {
		fmt.Println(file.Name())
		fileNames = append(fileNames, file.Name())
	}

	return fileNames

}
