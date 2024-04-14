package presenter

import "github.com/yosuke7040/dine-out-discoveries/usecase"

type ScrapingRestaurantPresenter struct{}

func NewScrapingRestaurantPresenter() usecase.ScrapingRestaurantPresenter {
	return &ScrapingRestaurantPresenter{}
}

func (p *ScrapingRestaurantPresenter) Output() usecase.ScrapingRestaurantOutput {
	return usecase.ScrapingRestaurantOutput{}
}
