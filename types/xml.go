package types

import "encoding/xml"

type Document struct {
	XMLName xml.Name `xml:"document"`
	Body    Body     `xml:"body"`
}
type Body struct {
	Paragraphs []Paragraph `xml:"p"`
}
type Paragraph struct {
	Runs []Run `xml:"r"`
}
type Run struct {
	Text *Text `xml:"t"`
	Sym  *Sym  `xml:"sym"`
}
type Text struct {
	Text string `xml:",chardata"`
}
type Sym struct {
	Font string `xml:"font,attr"`
	Char string `xml:"char,attr"`
}
