package logenios

import (
	"encoding/xml"
	"time"
)

type Time time.Time

func (t Time) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	dateString := time.Time(t).Format("2006-01-02 15:04:05")
	e.EncodeElement(dateString, start)

	return nil
}
