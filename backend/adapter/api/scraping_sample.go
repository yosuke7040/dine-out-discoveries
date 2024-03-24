package action

import (
	"log/slog"
	"net/http"

	"github.com/yosuke7040/dine-out-discoveries/adapter/response"
	"github.com/yosuke7040/dine-out-discoveries/usecase"
)

type ScrapingSampleAction struct {
	uc usecase.ScrapingSampleUseCase
}

func NewScrapingSampleAction(uc usecase.ScrapingSampleUseCase) *ScrapingSampleAction {
	return &ScrapingSampleAction{uc: uc}
}

func (a *ScrapingSampleAction) Execute(w http.ResponseWriter, r *http.Request) {
	var input usecase.ScrapingSampleInput

	// 仮
	input.Url = "https://www.tokyo-dome.co.jp/dome/event/schedule.html"

	// TODO: リクエスト受け付けてvalidationする

	// 具体的な処理はユースケースに任せる
	output, err := a.uc.Execute(r.Context(), input)
	if err != nil {
		a.handleErr(w, err)
		return
	}

	// ユースケースからのoutputをコントローラーに渡す
	response.NewSuccess(output, http.StatusOK).Send(w)
}

func (a *ScrapingSampleAction) handleErr(w http.ResponseWriter, err error) {
	slog.Error("Error scraping sample", err)
	response.NewError(err, http.StatusInternalServerError).Send(w)
	return
}
