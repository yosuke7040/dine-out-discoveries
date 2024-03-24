package presenter

import "github.com/yosuke7040/dine-out-discoveries/usecase"

type scrapingSamplePresenter struct{}

func NewScrapingSamplePresenter() usecase.ScrapingSamplePresenter {
	return &scrapingSamplePresenter{}
}

func (p *scrapingSamplePresenter) Output() usecase.ScrapingSampleOutput {
	return usecase.ScrapingSampleOutput{}
}
