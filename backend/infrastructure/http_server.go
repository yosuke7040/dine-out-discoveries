package infrastructure

import (
	"log/slog"
	"strconv"
	"time"

	"github.com/yosuke7040/dine-out-discoveries/infrastructure/router"
)

type config struct {
	webServer     router.Server
	webServerPort router.Port
	ctxTimeout    time.Duration
	// dbSQL repository.SQL
	// scraping adapter.CollyScraping
}

func NewConfig() *config {
	return &config{}
}

func (c *config) ContextTimeout(t time.Duration) *config {
	c.ctxTimeout = t
	return c
}

// func (c *config) ScrapingInstance() *config {
// 	scraping, err := scraping.NewScrapingFactory()
// 	if err != nil {
// 		slog.Error("Error configured scraping instance", "err", err)
// 	}
// 	c.scraping = scraping
// 	return c
// }

func (c *config) WebServer() *config {
	s, err := router.NewWebServerFactory(
		c.webServerPort,
		c.ctxTimeout,
		// c.scraping,
	)

	if err != nil {
		slog.Error("Error configured router server", err)
	}

	slog.Info("Successfully configured router server")

	c.webServer = s
	return c
}

func (c *config) WebServerPort(port string) *config {
	p, err := strconv.ParseInt(port, 10, 64)
	if err != nil {
		slog.Error("Error parsing port", err)
	}

	c.webServerPort = router.Port(p)
	return c
}

func (c *config) Start() {
	c.webServer.Listen()
}
