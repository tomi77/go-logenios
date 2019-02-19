package logenios

import (
	"encoding/xml"
	"time"
)

type Position struct {
	XMLName xml.Name `xml:"tem:Position"`

	CID       int32   `xml:"tem:CID"`
	ModemID   int64   `xml:"tem:ModemID"`
	Latitude  float32 `xml:"tem:Lat"`
	Longitude float32 `xml:"tem:Lon"`
	Time      Time    `xml:"tem:Time"`
}

func NewPosition(cID int32, modemID int64, latitude, longitude float32, time time.Time) *Position {
	return &Position{
		CID:       cID,
		ModemID:   modemID,
		Latitude:  latitude,
		Longitude: longitude,
		Time:      Time(time),
	}
}
