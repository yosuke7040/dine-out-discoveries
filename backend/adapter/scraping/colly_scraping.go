package scraping

import (
	"context"

	"github.com/gocolly/colly/v2"
)

type CollyScraping interface {
	// GetSample(context.Context, string) error
	OnHTML(context.Context, string, colly.HTMLCallback) error
	Visit(context.Context, string) error
	// OnRequest(context.Context, colly.HTMLCallback) error
}
