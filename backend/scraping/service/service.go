package service

import (
	"log/slog"
	"strconv"

	"github.com/gocolly/colly/v2"
)

type ServiceInterface interface {
	ScrapingTopPage(url string) error
	ScrapingReviewURLLists(url string) ([]string, error)
	ScrapingReviews(url string) error
}

type serviceStruct struct {
	// TODO: db
	cl *colly.Collector
}

func NewScrapingService(cl *colly.Collector) ServiceInterface {
	return &serviceStruct{cl: cl}
}

func (s *serviceStruct) ScrapingTopPage(url string) error {
	s.cl.OnHTML("section.rdheader-info-wrap", func(e *colly.HTMLElement) {
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

	s.cl.OnRequest(func(r *colly.Request) {
		slog.Info("ScrapingTopPage", "Visiting", r.URL)
	})

	s.cl.Visit(url)

	return nil
}

// https://tabelog.com/tokyo/A1305/A130501/13136428/dtlrvwlst/COND-0/smp1/D-like/1/?smp=1&lc=0&rvw_part=all
// D-like/1の部分がページ数
// 口コミのURLを一覧を取得
func (s *serviceStruct) ScrapingReviewURLLists(url string) ([]string, error) {
	var reviewURLs []string

	// 口コミのURLを一覧を取得
	s.cl.OnHTML("div.rstdtl-rvwlst", func(e *colly.HTMLElement) {
		e.ForEach("div.js-rvw-item-clickable-area", func(_ int, el *colly.HTMLElement) {
			link := el.ChildAttr("a.js-link-bookmark-detail", "data-detail-url")
			if link != "" {
				reviewURLs = append(reviewURLs, link)
			}
		})
	})

	s.cl.OnRequest(func(r *colly.Request) {
		slog.Info("ScrapingTopPage", "Visiting", r.URL)
	})

	s.cl.Visit(url)

	slog.Info("ScrapingReviews", "reviewURLs", reviewURLs)
	// fmt.Printf("----- reviewURLs: %#v\n", reviewURLs)
	return reviewURLs, nil
}

func (s *serviceStruct) ScrapingReviews(url string) {
	slog.Info("--- getReviewURLLists ----", "url", url)
}

// 口コミID、口コミ、口コミに対するいいね、スコアを取得
