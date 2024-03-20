package main

import (
	"context"
	"log"
	"net/http"
	"time"

	"github.com/gocolly/colly/v2"
	"github.com/yosuke7040/dine-out-discoveries/scraping/controller"
	"github.com/yosuke7040/dine-out-discoveries/scraping/service"
	"github.com/yosuke7040/dine-out-discoveries/scraping/usecase"
)

func main() {
	_, cancel := context.WithCancel(context.Background())
	defer cancel()

	c := colly.NewCollector()
	c.Limit(&colly.LimitRule{
		DomainGlob:  "*tabelog.com*",
		Delay:       1 * time.Second,
		RandomDelay: 1 * time.Second,
	})

	mux := http.NewServeMux()

	mux.HandleFunc("GET /", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello, World!"))
	})

	scrapingService := service.NewScrapingService(c)
	scrapingUseCase := usecase.NewScrapingUseCase(scrapingService)
	scrapingController := controller.NewScrapingController(scrapingUseCase)

	mux.HandleFunc("GET /scraping", func(w http.ResponseWriter, r *http.Request) {
		scrapingController.Scraping(w, r)
	})

	// go 1.22からGET指定出来そう
	mux.HandleFunc("GET /scraping/restaurant", func(w http.ResponseWriter, r *http.Request) {
		scrapingController.CollectRestaurantInfo(w, r)
	})

	log.Fatal(http.ListenAndServe(":8000", nil))
}
