package main

import "fmt"

type TZ int

const (
	HOUR TZ = 60 * 60
	UTC  TZ = 0 * HOUR
	EST  TZ = -5 * HOUR
	CST  TZ = -6 * HOUR
)

var timeZones = map[TZ]string{
	UTC: "Universal Greenwich time",
	EST: "Eastern Standard Time",
	CST: "Central Standard Time",
}

func (tz TZ) String() string {
	for name, zone := range timeZones {
		if tz == name {
			return zone
		}
	}
	return ""
}

func main() {
	fmt.Println(EST)
	fmt.Println(0 * HOUR)
	fmt.Println(-6 * HOUR)
}
