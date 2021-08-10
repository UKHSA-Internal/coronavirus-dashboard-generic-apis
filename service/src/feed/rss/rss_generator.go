package rss

import (
	"encoding/xml"
	"fmt"

	"generic_apis/feed"
)

type SelfLink struct {
	Href string `xml:"href,attr"`
	Rel  string `xml:"rel,attr"`
	Type string `xml:"type,attr"`
}

type Channel struct {
	XMLName        xml.Name        `xml:"channel"`
	Title          string          `xml:"title"`
	Category       string          `xml:"category,omitempty"`
	Description    string          `xml:"description"`
	Link           string          `xml:"link"`
	Copyright      string          `xml:"copyright,omitempty"`
	LastBuildDate  string          `xml:"lastBuildDate,omitempty"`
	PubDate        string          `xml:"pubDate,omitempty"`
	Ttl            int             `xml:"ttl,omitempty"`
	WebMaster      string          `xml:"webMaster,omitempty"`
	ManagingEditor string          `xml:"managingEditor,omitempty"`
	Language       string          `xml:"language,omitempty"`
	Generator      string          `xml:"generator,omitempty"`
	SelfLink       *SelfLink       `xml:"atom:link"`
	Items          *[]feed.Payload `xml:"item"`
}

const XmlHeader = xml.Header +
	`<?xml-stylesheet type="text/xsl" href="https://api.coronavirus.data.gov.uk/generic/xsl/rss.xsl"?>
<rss version="2.0" xmlns:atom="http://www.w3.org/2005/Atom">
%s
</rss>`

func (channel *Channel) GenerateFeed(components *feed.Components) ([]byte, error) {

	channel.Title = "UK Coronavirus Dashboard - RSS feed"

	channel.Description = fmt.Sprintf(
		`RSS feed of '%s' released in relation to the UK Coronavirus Dashboard.`,
		components.Category,
	)
	channel.Category = components.Category
	channel.Link = "https://coronavirus.data.gov.uk/"
	channel.Copyright = "2021 - Public Health England. Open Government License."
	channel.ManagingEditor = "coronavirus-tracker@phe.gov.uk (Coronavirus Dashboard Team)"
	channel.Ttl = 300
	channel.Language = "en-gb"
	channel.Generator = "UK Coronavirus Dashboard - Generic API Service"
	timestamp := components.Timestamp.Format("02 Jan 2006 15:04 -0700")
	channel.LastBuildDate = timestamp
	channel.PubDate = timestamp
	channel.SelfLink = &SelfLink{
		Href: "https://api.coronavirus.data.gov.uk" + components.Endpoint,
		Rel:  "self",
		Type: "application/rss+xml",
	}
	channel.Items = components.Payload

	encoded, err := xml.MarshalIndent(channel, "  ", "  ")
	if err != nil {
		return nil, err
	}

	return []byte(fmt.Sprintf(XmlHeader, encoded)), nil

} // generateFeed
