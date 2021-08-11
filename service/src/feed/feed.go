package feed

import (
	"time"
)

type Guid struct {
	IsPermaLink string `xml:"isPermaLink,attr"`
	Guid        string `xml:",chardata"`
}

type Payload struct {
	Title       string    `xml:"title"`
	Description string    `xml:"description"`
	Link        string    `xml:"link"`
	Guid        *Guid     `xml:"guid"`
	Date        time.Time `xml:"dc:date"`
}

type Components struct {
	Payload         *[]Payload
	Timestamp       *time.Time
	Category        string
	Endpoint        string
	ApiEndpoint     string
	WebsiteEndpoint string
}
