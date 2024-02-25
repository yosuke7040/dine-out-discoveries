package main

import (
	"log"
	"net/http"

	"github.com/yosuke7040/dine-out-discoveries/scraping/controller"
	"github.com/yosuke7040/dine-out-discoveries/scraping/service"
	"github.com/yosuke7040/dine-out-discoveries/scraping/usecase"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello, World!"))
	})

	scrapingService := service.NewScrapingService()
	scrapingUseCase := usecase.NewScrapingUseCase(scrapingService)
	scrapingController := controller.NewScrapingController(scrapingUseCase)

	http.HandleFunc("/scraping", func(w http.ResponseWriter, r *http.Request) {
		scrapingController.Scraping(w, r)
	})
	http.HandleFunc("/scraping/collect", func(w http.ResponseWriter, r *http.Request) {
		scrapingController.CollectRestaurantInfo(w, r)
	})

	log.Fatal(http.ListenAndServe(":8000", nil))
}
