package snr

import (
	"encoding/xml"
	"time"
)

// XMLFraga struct
type XMLFraga struct {
}

// XMLItem struct
type XMLItem struct {
	XMLName xml.Name `xml:",name"`
	XMLNS   string   `xml:"xmlns:com,attr"`
	Data    string   `xml:",chardata"`
}

// Informationshuvud struct
type Informationshuvud struct {
	Avsandare          XMLItem
	Sekvensnummer      XMLItem
	MeddelandeTyp      XMLItem
	Referens           XMLItem
	MeddelandeDatumTid XMLItem
	SessionsID         XMLItem
	Referens2          XMLItem
}

// Informationshuvud200 struct
type Informationshuvud200 struct {
	Avsandare          string `xml:"inf:Avsandare"`
	Sekvensnummer      string `xml:"inf:Sekvensnummer"`
	MeddelandeTyp      string `xml:"inf:MeddelandeTyp"`
	Referens           string `xml:"inf:Referens"`
	MeddelandeDatumTid string `xml:"inf:MeddelandeDatumTid"`
	SessionsID         string `xml:"inf:SessionsId"`
	Referens2          string `xml:"inf:Referens2"`
}

const (
	xmlns = "http://www.bolagsverket.se/schemas/common"
)

// NewInformationshuvud200 func
func NewInformationshuvud200() *Informationshuvud200 {
	t := time.Now()
	tf := t.Format(time.RFC3339Nano)
	return &Informationshuvud200{
		Avsandare:          "String",
		Sekvensnummer:      "String",
		MeddelandeTyp:      "String",
		Referens:           "String",
		MeddelandeDatumTid: tf,
		SessionsID:         "String",
		Referens2:          "String",
	}
}

// NewInformationshuvud func
func NewInformationshuvud() *Informationshuvud {
	t := time.Now()
	tf := t.Format(time.RFC3339Nano)
	return &Informationshuvud{
		Avsandare: XMLItem{
			XMLName: xml.Name{Local: "com:Avsandare"},
			XMLNS:   xmlns,
			Data:    "String",
		},
		Sekvensnummer: XMLItem{
			XMLName: xml.Name{Local: "com:Sekvensnummer"},
			XMLNS:   xmlns,
			Data:    "String",
		},
		MeddelandeTyp: XMLItem{
			XMLName: xml.Name{Local: "com:MeddelandeTyp"},
			XMLNS:   xmlns,
			Data:    "String",
		},
		Referens: XMLItem{
			XMLName: xml.Name{Local: "com:Referens"},
			XMLNS:   xmlns,
			Data:    "String",
		},
		MeddelandeDatumTid: XMLItem{
			XMLName: xml.Name{Local: "com:MeddelandeDatumTid"},
			XMLNS:   xmlns,
			Data:    tf,
		},
		SessionsID: XMLItem{
			XMLName: xml.Name{Local: "com:SessionsId"},
			XMLNS:   xmlns,
			Data:    "String",
		},
		Referens2: XMLItem{
			XMLName: xml.Name{Local: "com:Referens2"},
			XMLNS:   xmlns,
			Data:    "String",
		},
	}
}
