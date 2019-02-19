package logenios

import (
	"encoding/xml"
)

type AddPosition struct {
	XMLName xml.Name `xml:"soapenv:Envelope"`

	SoapEnv string `xml:"xmlns:soapenv,attr"`
	Tem     string `xml:"xmlns:tem,attr"`

	Header Header

	Customer string `xml:"soapenv:Body>tem:AddPosition>tem:Customer"`
	User     string `xml:"soapenv:Body>tem:AddPosition>tem:User"`
	Password string `xml:"soapenv:Body>tem:AddPosition>tem:Password"`

	Positions Positions `xml:"soapenv:Body>tem:AddPosition>tem:positions"`
}

func (p *AddPosition) AddPosition(position Position) {
	p.Positions.Positions = append(p.Positions.Positions, position)
}

// NewAddPosition creates a new AddPosition struct and init it
func NewAddPosition(customer, user, password string) *AddPosition {
	return &AddPosition{
		SoapEnv:   "http://schemas.xmlsoap.org/soap/envelope/",
		Tem:       "http://tempuri.org/",
		Customer:  customer,
		User:      user,
		Password:  password,
		Positions: Positions{},
	}
}
