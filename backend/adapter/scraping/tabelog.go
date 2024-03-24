package scraping

import (
	"context"
	"log/slog"
	"time"

	"github.com/gocolly/colly/v2"
)

type TabelogScraping struct {
	scraping CollyScraping
}

type TabelogInterface interface {
	GetSample(context.Context, string) error
}

func NewTabelogScraping(scraping CollyScraping) *TabelogScraping {
	return &TabelogScraping{scraping: scraping}
}

func (s *TabelogScraping) GetSample(ctx context.Context, url string) error {

	jst, _ := time.LoadLocation("Asia/Tokyo")
	url = "https://www.tokyo-dome.co.jp/dome/event/schedule.html"

	selector := "div.c-mod-tab__body:nth-child(2) > table > tbody"
	innerSelector := "tr.c-mod-calender__item"
	dateSelector := "th > span:nth-child(1)"
	categorySelector := "td:nth-child(2) > div > div:nth-child(1) > p > span"
	titleSelector := "td > div > div:nth-child(2) > p.c-mod-calender__links"
	timeSelector := "td > div > div:nth-child(2) > p:nth-child(2)"

	var event string
	// TODO: ここでcolly使ったらadapterの責務を超えているのでは？
	s.scraping.OnHTML(ctx, selector, func(e *colly.HTMLElement) {
		e.ForEach(innerSelector, func(_ int, s *colly.HTMLElement) {
			date := s.ChildText(dateSelector)
			category := s.ChildText(categorySelector)
			title := s.ChildText(titleSelector)
			info := s.ChildText(timeSelector)
			today := time.Now().In(jst).Format("02")

			if date == today {
				if title == "" {
					event = "イベントなし"
				} else {
					event = title + "（" + category + "）" + "\n" + info
				}
			}
		})
	})

	s.scraping.Visit(ctx, url)

	slog.Info("usecase scraping....")
	slog.Info(event)
	slog.Info("usecase scraping end....")

	return nil
}
