package router

import (
	"time"
)

type Server interface {
	Listen()
}

type Port uint16

func NewWebServerFactory(
	port Port,
	ctxTimeout time.Duration,
	// scraping adapter.CollyScraping,
) (Server, error) {
	return newServerMuxEngine(port, ctxTimeout), nil
}
