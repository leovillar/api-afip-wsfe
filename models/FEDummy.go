package models

import "encoding/xml"

type ReqFEDummy struct {
	XMLName xml.Name `xml:"soapenv:Envelope"`
	Soapenv string   `xml:"xmlns:soapenv,attr"`
	Ar      string   `xml:"xmlns:ar,attr"`
	Header  string   `xml:"soapenv:Header"`
	Body    struct {
		FEDummy string `xml:"ar:FEDummy"`
	} `xml:"soapenv:Body"`
}

type RespFEDummy struct {
	XMLName xml.Name `xml:"Envelope"`
	//Xsi     string   `xml:"xsi,attr"`
	//Xsd     string   `xml:"xsd,attr"`
	//Soap    string   `xml:"soap,attr"`
	Body struct {
		FEDummyResponse struct {
			//Xmlns         string `xml:"xmlns,attr"`
			FEDummyResult struct {
				AppServer  string `xml:"AppServer" json:"appServer"`
				DbServer   string `xml:"DbServer" json:"dbServer"`
				AuthServer string `xml:"AuthServer" json:"authServer"`
			} `xml:"FEDummyResult" json:"feDummyResult"`
		} `xml:"FEDummyResponse" json:"feDummyResponse"`
	} `xml:"Body" json:"body"`
}

func (r ReqFEDummy) NewFEDummy() []byte {
	r.Soapenv = "http://schemas.xmlsoap.org/soap/envelope/"
	r.Ar = "http://ar.gov.afip.dif.FEV1/"
	output, _ := xml.MarshalIndent(r, "  ", "    ")
	return output
}
