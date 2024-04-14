package action

import (
	"log/slog"
	"net/http"

	"github.com/yosuke7040/dine-out-discoveries/adapter/response"
	"github.com/yosuke7040/dine-out-discoveries/usecase"
)

type ScrapingRestaurantAction struct {
	uc usecase.ScrapingRestaurantUseCase
}

func NewScrapingRestaurantAction(uc usecase.ScrapingRestaurantUseCase) *ScrapingRestaurantAction {
	return &ScrapingRestaurantAction{uc: uc}
}

func (a *ScrapingRestaurantAction) Execute(w http.ResponseWriter, r *http.Request) {
	var input usecase.ScrapingRestaurantInput

	// TODO: 仮
	input.Url = "https://tabelog.com/tokyo/A1305/A130501/13136428/"

	output, err := a.uc.Execute(r.Context(), input)
	if err != nil {
		a.handleErr(w, err)
		return
	}

	response.NewSuccess(output, http.StatusOK).Send(w)
}

func (a *ScrapingRestaurantAction) handleErr(w http.ResponseWriter, err error) {
	slog.Error("Error scraping restaurant", err)
	response.NewError(err, http.StatusInternalServerError).Send(w)

	return
}
