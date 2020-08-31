package pixie

import "encoding/xml"

//Response represents the XML response element
type Response struct {
	XMLName     xml.Name `xml:"response"`
	Code        string   `xml:"code,attr"`
	Description string   `xml:"description,attr"`
}
