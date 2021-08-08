package atom

import (
	"encoding/xml"
	"fmt"
	"time"

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
	Rel  string `xml:"rel,attr"`
	Href string `xml:"href,attr"`
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
	Link    *Link    `xml:"link"`
	Id      string   `xml:"id"`
	Author  string   `xml:"author>name"`
}

type Feed struct {
	XMLName   xml.Name   `xml:"feed"`
	Xmlns     string     `xml:"xmlns,attr"`
	Title     string     `xml:"title"`
	Category  *Category  `xml:"category,omitempty"`
	Link      *Link      `xml:"link"`
	Rights    string     `xml:"rights,omitempty"`
	Updated   string     `xml:"updated"`
	Id        string     `xml:"id"`
	Generator *Generator `xml:"generator,omitempty"`
	Items     *[]Payload `xml:"entry"`
}

const XhtmlWrapper = `<div xmlns="http://www.w3.org/1999/xhtml">%s</div>`

func (feed *Feed) GenerateFeed(payload *[]feed.Payload, timestamp *time.Time) ([]byte, error) {

	feed.Title = "UK Coronavirus Dashboard - Atom feed"
	feed.Xmlns = "http://www.w3.org/2005/Atom"
	feed.Category = &Category{Term: "Announcements"}
	feed.Link = &Link{Rel: "self", Href: "https://api.coronavirus.data.gov.uk/generic/announcement/atom.xml"}
	feed.Id = "https://coronavirus.data.gov.uk/"
	feed.Rights = "2021 - Public Health England. Open Government License."
	feed.Generator = &Generator{
		Uri:       "https://api.coronavirus.data.gov.uk/generic/announcement/atom.xml",
		Version:   1,
		Generator: "UK Coronavirus Dashboard - Generic API Service",
	}
	feed.Updated = timestamp.Format("2006-01-02T15:04:05Z")

	atomPayload := make([]Payload, len(*payload))

	mdOpts := html.RendererOptions{
		Flags: html.UseXHTML |
			html.HrefTargetBlank |
			html.NoreferrerLinks |
			html.NoopenerLinks,
	}
	mdRenderer := html.NewRenderer(mdOpts)

	for index, item := range *payload {
		md := []byte(item.Description)
		atomPayload[index].Content = &Content{
			Type:    "xhtml",
			Content: fmt.Sprintf(XhtmlWrapper, markdown.ToHTML(md, nil, mdRenderer)),
		}

		lastUpdate, _ := time.Parse("Mon, 02 Jan 2006 15:04:05 -0700", item.PubDate)
		atomPayload[index].Updated = lastUpdate.Format("2006-01-02T15:04:05Z")
		atomPayload[index].Link = &Link{Rel: "alternate", Href: item.Link}
		atomPayload[index].Id = "urn:uuid:" + item.Guid.Guid
		atomPayload[index].Title = item.Title
		atomPayload[index].Author = "UK Coronavirus Dashboard Team"
	}
	feed.Items = &atomPayload

	encoded, err := xml.MarshalIndent(feed, "  ", "  ")
	if err != nil {
		return nil, err
	}

	return []byte(xml.Header + string(encoded)), nil

} // generateFeed