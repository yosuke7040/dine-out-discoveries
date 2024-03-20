package usecase

import (
	"context"
	"log/slog"
	"strconv"
	"time"

	"github.com/gocolly/colly/v2"
	"github.com/yosuke7040/dine-out-discoveries/scraping/service"
)

type UsecaseInterface interface {
	Scraping()
	CollectRestaurantInfo(ctx context.Context, in string)
}

// TODO: 引数はinとかでまとめるほうが良いかも
// type CollectRestaurantInfoRequest struct {
// 	urls []string
// }

type usecaseStruct struct {
	ss service.ServiceInterface
}

func NewScrapingUseCase(ss service.ServiceInterface) UsecaseInterface {
	return &usecaseStruct{ss: ss}
}

// スクレイピングお試し用
func (u *usecaseStruct) Scraping() {
	jst, _ := time.LoadLocation("Asia/Tokyo")
	url := "https://www.tokyo-dome.co.jp/dome/event/schedule.html"

	c := colly.NewCollector()

	selector := "div.c-mod-tab__body:nth-child(2) > table > tbody"
	innerSelector := "tr.c-mod-calender__item"
	dateSelector := "th > span:nth-child(1)"
	categorySelector := "td:nth-child(2) > div > div:nth-child(1) > p > span"
	titleSelector := "td > div > div:nth-child(2) > p.c-mod-calender__links"
	timeSelector := "td > div > div:nth-child(2) > p:nth-child(2)"

	var event string
	c.OnHTML(selector, func(e *colly.HTMLElement) {
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

	c.Visit(url)

	slog.Info("usecase scraping....")
	slog.Info(event)
	slog.Info("usecase scraping end....")
}

// NOTE: URLについて
// https://~~/dtlmenu/ -> メニュー・コース
// https://~~/dtlphotolst/smp2/ -> 写真
// https://~~/dtlrvwlst -> 口コミ
// https://~~/dtlmap -> 地図
func (u *usecaseStruct) CollectRestaurantInfo(ctx context.Context, url string) {
	// topページの情報を取得
	if err := u.ss.ScrapingTopPage(url); err != nil {
		slog.Error("ScrapingTopPage: ", err)
	}

	// 口コミの情報を20件ずつ取得
	// httpS://~~~com/~~/dtlrvwlst/COND-0/smp1/?smp=1&lc=0&rvw_part=all&PG=1
	// ↓いいね順の口コミ
	// https:/〜〜/dtlrvwlst/COND-0/smp1/D-like/?smp=1&lc=0&rvw_part=all&pg=1
	for i := 0; i < 2; i++ {
		reviewURLs, err := u.ss.ScrapingReviewURLLists(url + "/dtlrvwlst/COND-0/smp1/D-like/" + strconv.Itoa(i+1) + "/?smp=1&lc=0&rvw_part=all")
		if err != nil {
			slog.Error("")
		}

		u.ss.
	}
}
