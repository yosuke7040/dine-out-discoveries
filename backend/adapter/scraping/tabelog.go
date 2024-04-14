package scraping

import (
	"context"
	"log/slog"
	"strconv"
	"time"

	"github.com/gocolly/colly/v2"
)

type TabelogScraping struct {
	// scraping CollyScraping
}

type TabelogInterface interface {
	GetSample(context.Context, string) error
	GetRestaurantTopPage(context.Context, string) error
}

func NewTabelogScraping() *TabelogScraping {
	return &TabelogScraping{}
}

// func NewTabelogScraping(scraping CollyScraping) *TabelogScraping {
// 	return &TabelogScraping{scraping: scraping}
// }

func (s *TabelogScraping) GetSample(ctx context.Context, url string) error {
	slog.Info("usecase scraping....")

	cl := colly.NewCollector(
	// colly.Debugger(&debug.LogDebugger{}),
	// 非同期か並行にしたい場合はCollector.Wait()を呼ぶ
	// colly.Async(),
	)
	// cl.Limit(&colly.LimitRule{
	// 	DomainGlob:  "*tabelog.com*",
	// 	Delay:       1 * time.Second,
	// 	RandomDelay: 1 * time.Second,
	// })

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
	cl.OnHTML(selector, func(e *colly.HTMLElement) {
		e.ForEach(innerSelector, func(_ int, s *colly.HTMLElement) {
			date := s.ChildText(dateSelector)
			category := s.ChildText(categorySelector)
			title := s.ChildText(titleSelector)
			info := s.ChildText(timeSelector)
			today := time.Now().In(jst).Format("02")

			slog.Info("colly scraping....", "date", date, "category", category, "title", title, "info", info, "today", today)

			if date == today {
				if title == "" {
					event = "イベントなし"
				} else {
					event = title + "（" + category + "）" + "\n" + info
				}
			}
		})
	})

	cl.OnRequest(func(r *colly.Request) {
		slog.Info("Visiting", "url", r.URL.String())
	})

	err := cl.Visit(url)
	// err := s.scraping.Visit(ctx, url)
	if err != nil {
		slog.Error("Error scraping", "err", err)
		return err
	}
	slog.Info("event: ", "event", event)
	slog.Info("usecase scraping end...")

	return nil
}

func (s *TabelogScraping) GetRestaurantTopPage(ctx context.Context, url string) error {
	slog.Info("usecase scraping....")

	cl := colly.NewCollector()
	cl.Limit(&colly.LimitRule{
		DomainGlob:  "*tabelog.com*",
		Delay:       1 * time.Second,
		RandomDelay: 1 * time.Second,
	})

	cl.OnHTML("section.rdheader-info-wrap", func(e *colly.HTMLElement) {
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

	cl.OnRequest(func(r *colly.Request) {
		slog.Info("ScrapingTopPage", "Visiting", r.URL)
	})

	cl.Visit(url)

	return nil
}

func (s *TabelogScraping) GetReviewUrlLists(ctx context.Context, url string) ([]string, error) {
	cl := colly.NewCollector()
	cl.Limit(&colly.LimitRule{
		DomainGlob:  "*tabelog.com*",
		Delay:       1 * time.Second,
		RandomDelay: 1 * time.Second,
	})

	var reviewURLs []string

	// 口コミのURLを一覧を取得
	cl.OnHTML("div.rstdtl-rvwlst", func(e *colly.HTMLElement) {
		e.ForEach("div.js-rvw-item-clickable-area", func(_ int, el *colly.HTMLElement) {
			link := el.ChildAttr("a.js-link-bookmark-detail", "data-detail-url")
			if link != "" {
				reviewURLs = append(reviewURLs, link)
			}
		})
	})

	cl.OnRequest(func(r *colly.Request) {
		slog.Info("ScrapingTopPage", "Visiting", r.URL)
	})

	cl.Visit(url)

	slog.Info("ScrapingReviews", "reviewURLs", reviewURLs)
	// fmt.Printf("----- reviewURLs: %#v\n", reviewURLs)
	return reviewURLs, nil
}
