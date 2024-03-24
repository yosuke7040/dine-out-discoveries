package scraping

import (
	"context"

	"github.com/gocolly/colly/v2"
)

type collyHandler struct {
	cl *colly.Collector
}

func NewCollyHandler(c *config) (*collyHandler, error) {
	cl := colly.NewCollector()
	cl.Limit(&colly.LimitRule{
		DomainGlob:  c.DomainGlob,
		Delay:       c.Delay,
		RandomDelay: c.RandomDelay,
	})
	return &collyHandler{cl: cl}, nil
}

func (h *collyHandler) OnHTML(ctx context.Context, goquerySelector string, f colly.HTMLCallback) error {
	h.cl.OnHTML(goquerySelector, f)
	return nil
}

// func (h *collyHandler) OnRequest(ctx context.Context, f colly.RequestCallback) error {
// 	h.cl.OnRequest(f)
// 	return nil
// }

func (h *collyHandler) Visit(ctx context.Context, url string) error {
	h.cl.Visit(url)
	return nil
}
