package logenios

import "encoding/xml"

type Header struct {
	XMLName xml.Name `xml:"soapenv:Header"`
}
