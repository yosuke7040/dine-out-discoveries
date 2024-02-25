package service

import (
	"log/slog"
	"strconv"

	"github.com/gocolly/colly"
)

type ServiceInterface interface {
	ScrapingTopPage(url string) error
	ScrapingReviews(url string) error
}

type serviceStruct struct {
	// db
}

func NewScrapingService() ServiceInterface {
	return &serviceStruct{}
}

func (s *serviceStruct) ScrapingTopPage(url string) error {
	c := colly.NewCollector()

	c.OnHTML("section.rdheader-info-wrap", func(e *colly.HTMLElement) {
		storeName := e.ChildText("div.rdheader-rstname > h2.display-name")
		aliasStoreName := e.ChildText("div.rdheader-rstname > span.alias")
		scoreStr := e.ChildText("b.c-rating__val.rdheader-rating__score-val > span.rdheader-rating__score-val-dtl")
		score, err := strconv.ParseFloat(scoreStr, 32)
		if err != nil {
			slog.Error("ScrapingTopPage", "Score conversion error: ", err)
			return
		}

		// NOTE: 予算関連もレコメンドで使うなら必要になる
		// 一旦、昼と夜が一緒になってるから注意
		// values := e.ChildText("span.c-rating-v3__val > a.rdheader-budget__price-target")

		// TODO: DBへ
		slog.Info("ScrapingTopPage", "storeName", storeName, "aliasStoreName", aliasStoreName, "score", score)
	})

	c.OnRequest(func(r *colly.Request) {
		slog.Info("ScrapingTopPage", "Visiting", r.URL)
	})

	c.Visit(url)

	return nil
}

func (s *serviceStruct) ScrapingReviews(url string) error {

	return nil
}
