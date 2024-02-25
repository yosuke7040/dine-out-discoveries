package usecase

import (
	"context"
	"log/slog"

	// "fmt"
	"time"

	"github.com/gocolly/colly"
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
	err := u.ss.ScrapingTopPage(url)
	if err != nil {
		slog.Error("ScrapingTopPage: ", err)
		// fmt.Println("ScrapingTopPage error: ")
	}

	// 口コミの情報を取得
	err = u.ss.ScrapingReviews(url + "dtlrvwlst")
	if err != nil {
		slog.Error("")
	}
}
