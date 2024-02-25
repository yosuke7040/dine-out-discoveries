package controller

import (
	"net/http"

	"github.com/yosuke7040/dine-out-discoveries/scraping/usecase"
)

type controllerInterface interface {
	Scraping(w http.ResponseWriter, r *http.Request)
	CollectRestaurantInfo(w http.ResponseWriter, r *http.Request)
}

type controllerStruct struct {
	su usecase.UsecaseInterface
}

func NewScrapingController(su usecase.UsecaseInterface) controllerInterface {
	return &controllerStruct{su: su}
}

// スクレイピングお試し用
func (c *controllerStruct) Scraping(w http.ResponseWriter, r *http.Request) {
	c.su.Scraping()
	w.Write([]byte("scraping..."))
}

func (c *controllerStruct) CollectRestaurantInfo(w http.ResponseWriter, r *http.Request) {
	url := "https://tabelog.com/tokyo/A1305/A130501/13136428/"
	c.su.CollectRestaurantInfo(r.Context(), url)

	w.WriteHeader(http.StatusAccepted)
	w.Write([]byte("CollectRestaurantInfo..."))
}
