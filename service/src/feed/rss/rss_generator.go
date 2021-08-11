package rss

import (
	"encoding/xml"
	"fmt"
	"strings"

	"generic_apis/feed"
	"github.com/gomarkdown/markdown"
	"github.com/gomarkdown/markdown/html"
)

type Payload struct {
	Title       string     `xml:"title"`
	Description string     `xml:"description"`
	Content     *Content   `xml:"content:encoded"`
	Link        string     `xml:"link"`
	Guid        *feed.Guid `xml:"guid"`
	Date        string     `xml:"dc:date"`
	Subject     string     `xml:"dc:subject,omitempty"`
}

type Content struct {
	Data string `xml:",innerxml"`
}

type SelfLink struct {
	Href string `xml:"href,attr"`
	Rel  string `xml:"rel,attr"`
	Type string `xml:"type,attr"`
}

type Image struct {
	Url         string `xml:"url"`
	Title       string `xml:"title"`
	Link        string `xml:"link"`
	Width       string `xml:"width,omitempty"`
	Height      string `xml:"height,omitempty"`
	Description string `xml:"description,omitempty"`
}

type AdminAttrs struct {
	Resource string `xml:"rdf:resource,attr,omitempty"`
}

type Channel struct {
	XMLName     xml.Name `xml:"channel"`
	Title       string   `xml:"title"`
	Category    string   `xml:"category,omitempty"`
	Description string   `xml:"description"`
	Link        string   `xml:"link"`
	Image       *Image   `xml:"image"`

	Rights string `xml:"dc:rights,omitempty"`

	LastBuildDate string `xml:"lastBuildDate,omitempty"`
	Date          string `xml:"dc:date,omitempty"`
	Ttl           int    `xml:"ttl,omitempty"`
	UpdatePeriod  string `xml:"sy:updatePeriod,omitempty"`
	UpdateFreq    int    `xml:"sy:updateFrequency,omitempty"`
	UpdateBase    string `xml:"sy:updateBase,omitempty"`

	WebMaster   string      `xml:"webMaster,omitempty"`
	Creator     string      `xml:"dc:creator,omitempty"`
	ErrorReport *AdminAttrs `xml:"admin:errorReportsTo"`

	Language  string     `xml:"dc:language,omitempty"`
	Generator string     `xml:"generator,omitempty"`
	SelfLink  *SelfLink  `xml:"atom:link"`
	Items     *[]Payload `xml:"item"`
}

const XmlHeader = xml.Header +
	`<?xml-stylesheet type="text/xsl" href="https://api.coronavirus.data.gov.uk/generic/xsl/rss.xsl"?>
<rss version="2.0"
     xmlns:atom="http://www.w3.org/2005/Atom"
     xmlns:dc="http://purl.org/dc/elements/1.1/"
     xmlns:sy="http://purl.org/rss/1.0/modules/syndication/"
     xmlns:admin="http://webns.net/mvcb/"
     xmlns:rdf="http://www.w3.org/1999/02/22-rdf-syntax-ns#"
     xmlns:content="http://purl.org/rss/1.0/modules/content/">
%s
</rss>`

const xhtmlWrapper = "<![CDATA[%s]]>"

func (channel *Channel) GenerateFeed(components *feed.Components) ([]byte, error) {

	channel.Title = "UK Coronavirus Dashboard - RSS feed"

	channel.Description = fmt.Sprintf(
		`RSS feed of '%s' released in relation to the UK Coronavirus Dashboard.`,
		components.Category,
	)
	channel.Category = components.Category
	channel.Link = "https://coronavirus.data.gov.uk/details/announcements"
	channel.Rights = "2021 - Public Health England. Open Government License."
	channel.Image = &Image{
		Url:   "https://coronavirus.data.gov.uk/favicon.png",
		Title: "UK Coronavirus Dashboard - RSS feed",
		Link:  "https://coronavirus.data.gov.uk/details/announcements",
	}
	channel.Ttl = 300
	channel.Language = "en-gb"
	channel.Generator = "UK Coronavirus Dashboard - Generic API Service"
	channel.LastBuildDate = components.Timestamp.Format("02 Jan 2006 15:04 -0700")

	channel.Date = components.Timestamp.Format("2006-01-02T15:04:05-07:00")
	channel.UpdatePeriod = "hourly"
	channel.UpdateFreq = 1
	channel.UpdateBase = "2019-04-14T00:00:00+00:00"

	channel.SelfLink = &SelfLink{
		Href: "https://api.coronavirus.data.gov.uk" + components.Endpoint,
		Rel:  "self",
		Type: "application/rss+xml",
	}
	channel.Creator = "coronavirus-tracker@phe.gov.uk"
	channel.ErrorReport = &AdminAttrs{Resource: channel.Creator}

	rssPayload := make([]Payload, len(*components.Payload))

	mdOpts := html.RendererOptions{
		Flags: html.UseXHTML |
			html.HrefTargetBlank |
			html.NoreferrerLinks |
			html.NoopenerLinks,
	}
	mdRenderer := html.NewRenderer(mdOpts)

	for index, item := range *components.Payload {
		md := []byte(item.Description)

		rssPayload[index].Content = &Content{
			Data: strings.ReplaceAll(
				fmt.Sprintf(xhtmlWrapper, markdown.ToHTML(md, nil, mdRenderer)),
				`href="/`,
				`href="https://coronavirus.data.gov.uk/`,
			),
		}

		rssPayload[index].Description = item.Description
		rssPayload[index].Title = item.Title
		rssPayload[index].Link = item.Link
		rssPayload[index].Guid = item.Guid
		if item.Category != "" {
			rssPayload[index].Subject = item.Category
		}
		rssPayload[index].Date = item.Date.Format("2006-01-02T15:04:05-07:00")
	}

	channel.Items = &rssPayload

	encoded, err := xml.MarshalIndent(channel, "  ", "  ")
	if err != nil {
		return nil, err
	}

	return []byte(fmt.Sprintf(XmlHeader, encoded)), nil

} // generateFeed
