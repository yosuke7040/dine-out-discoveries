package scraping

// import (
// 	"context"

// 	"github.com/gocolly/colly/v2"
// 	"github.com/gocolly/colly/v2/debug"
// )

// // もしや、collectorは使い回すのではなく、都度生成するもの？？
// // 使い回すと2回目のOnHTMLが機能しない…キャッシュかなと思ったが、無効にするオプションがないのと
// // 下記URLでいくつかインスタンスを生成しているので都度生成するものかもしれない
// // https://github.com/gocolly/colly/blob/master/_examples/google_groups/google_groups.go
// type collyHandler struct {
// 	cl *colly.Collector
// }

// func NewCollyHandler(c *config) (*collyHandler, error) {
// 	cl := colly.NewCollector(
// 		colly.Debugger(&debug.LogDebugger{}),
// 		// 非同期か並行にしたい場合はCollector.Wait()を呼ぶ
// 		// colly.Async(),
// 	)
// 	cl.Limit(&colly.LimitRule{
// 		DomainGlob:  c.DomainGlob,
// 		Delay:       c.Delay,
// 		RandomDelay: c.RandomDelay,
// 	})
// 	return &collyHandler{cl: cl}, nil
// }

// func (h *collyHandler) OnHTML(ctx context.Context, goquerySelector string, f colly.HTMLCallback) error {
// 	h.cl.OnHTML(goquerySelector, f)
// 	return nil
// }

// func (h *collyHandler) OnRequest(ctx context.Context, f colly.RequestCallback) error {
// 	h.cl.OnRequest(f)
// 	return nil
// }

// func (h *collyHandler) Visit(ctx context.Context, url string) error {
// 	h.cl.Visit(url)
// 	return nil
// }
