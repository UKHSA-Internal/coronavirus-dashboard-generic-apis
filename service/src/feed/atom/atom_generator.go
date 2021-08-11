package atom

import (
	"encoding/xml"
	"fmt"
	"strings"

	"generic_apis/feed"
	"github.com/gomarkdown/markdown"
	"github.com/gomarkdown/markdown/html"
)

type Generator struct {
	Uri       string `xml:"uri,attr,omitempty"`
	Version   int    `xml:"version,attr,omitempty"`
	Generator string `xml:",chardata"`
}

type Link struct {
	Rel      string `xml:"rel,attr,omitempty"`
	Type     string `xml:"type,attr,omitempty"`
	Hreflang string `xml:"hreflang,attr,omitempty"`
	Title    string `xml:"title,attr,omitempty"`
	Href     string `xml:"href,attr"`
}

type Category struct {
	Term string `xml:"term,attr"`
}

type Content struct {
	Type    string `xml:"type,attr"`
	Content string `xml:",innerxml"`
}

type Source struct {
	Id      string `xml:"id"`
	Title   string `xml:"title,omitempty"`
	Updated string `xml:"updated,omitempty"`
}

type Payload struct {
	Title   string   `xml:"title"`
	Content *Content `xml:"content"`
	Updated string   `xml:"updated"`
	Link    *[]Link  `xml:"link"`
	Id      string   `xml:"id"`
	Author  string   `xml:"author>name"`
}

type Feed struct {
	XMLName   xml.Name   `xml:"feed"`
	Xmlns     string     `xml:"xmlns,attr"`
	Title     string     `xml:"title"`
	Icon      string     `xml:"icon"`
	Logo      string     `xml:"logo"`
	Category  *Category  `xml:"category,omitempty"`
	Link      *[]Link    `xml:"link"`
	Rights    string     `xml:"rights,omitempty"`
	Updated   string     `xml:"updated"`
	Id        string     `xml:"id"`
	Generator *Generator `xml:"generator,omitempty"`
	Items     *[]Payload `xml:"entry"`
}

const header = xml.Header +
	`<?xml-stylesheet type="text/xsl" href="https://api.coronavirus.data.gov.uk/generic/xsl/atom.xsl"?>` +
	"\n"

const xhtmlWrapper = `<div xmlns="http://www.w3.org/1999/xhtml">%s</div>`

func (feed *Feed) GenerateFeed(components *feed.Components) ([]byte, error) {

	feed.Title = "UK Coronavirus Dashboard - Atom feed"
	feed.Xmlns = "http://www.w3.org/2005/Atom"
	feed.Category = &Category{Term: components.Category}
	feed.Icon = "https://coronavirus.data.gov.uk/favicon.ico"
	feed.Logo = "https://coronavirus.data.gov.uk/favicon.png"
	feed.Link = &[]Link{
		{
			Rel:  "self",
			Href: "https://api.coronavirus.data.gov.uk" + components.Endpoint,
		},
		{
			Rel:  "alternate",
			Type: "text/html",
			Href: components.WebsiteEndpoint,
		},
	}
	feed.Id = components.WebsiteEndpoint
	feed.Rights = "2021 - Public Health England. Open Government License."
	feed.Generator = &Generator{
		Uri:       "https://api.coronavirus.data.gov.uk" + components.Endpoint,
		Version:   1,
		Generator: "UK Coronavirus Dashboard - Generic API Service",
	}
	feed.Updated = components.Timestamp.Format("2006-01-02T15:04:05-07:00")

	atomPayload := make([]Payload, len(*components.Payload))

	mdOpts := html.RendererOptions{
		Flags: html.UseXHTML |
			html.HrefTargetBlank |
			html.NoreferrerLinks |
			html.NoopenerLinks,
	}
	mdRenderer := html.NewRenderer(mdOpts)

	for index, item := range *components.Payload {
		md := []byte(item.Description)
		atomPayload[index].Content = &Content{
			Type: "xhtml",
			Content: strings.ReplaceAll(
				fmt.Sprintf(xhtmlWrapper, markdown.ToHTML(md, nil, mdRenderer)),
				`href="/`,
				`href="https://coronavirus.data.gov.uk/`,
			),
		}

		atomPayload[index].Updated = item.Date.Format("2006-01-02T15:04:05-07:00")
		atomPayload[index].Link = &[]Link{
			{
				Rel:      "alternate",
				Type:     "text/html",
				Href:     item.Link,
				Hreflang: "en-gb",
				Title:    "View record on the website",
			},
			{
				Rel:      "alternate",
				Type:     "application/json",
				Hreflang: "en-gb",
				Href:     components.ApiEndpoint + item.Guid.Guid,
				Title:    "View JSON payload via the API",
			},
		}
		atomPayload[index].Id = "urn:uuid:" + item.Guid.Guid
		atomPayload[index].Title = item.Title
		atomPayload[index].Author = "UK Coronavirus Dashboard Team"
	}
	feed.Items = &atomPayload

	encoded, err := xml.MarshalIndent(feed, "  ", "  ")
	if err != nil {
		return nil, err
	}

	return []byte(header + string(encoded)), nil

} // generateFeed
