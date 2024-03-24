package router

import (
	"time"

	adapter "github.com/yosuke7040/dine-out-discoveries/adapter/scraping"
)

type Server interface {
	Listen()
}

type Port uint16

func NewWebServerFactory(
	port Port,
	ctxTimeout time.Duration,
	scraping adapter.CollyScraping,
) (Server, error) {
	return newServerMuxEngine(port, ctxTimeout, scraping), nil
}
