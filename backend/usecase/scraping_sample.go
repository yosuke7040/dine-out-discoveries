package usecase

import (
	"context"
	"time"

	adapter "github.com/yosuke7040/dine-out-discoveries/adapter/scraping"
)

type (
	ScrapingSampleUseCase interface {
		Execute(context.Context, ScrapingSampleInput) (ScrapingSampleOutput, error)
	}

	ScrapingSampleInput struct {
		Url string `json:"url"`
	}

	ScrapingSamplePresenter interface {
		Output() ScrapingSampleOutput
	}

	ScrapingSampleOutput struct {
		Result string `json:"id"`
	}

	ScrapingSampleInteractor struct {
		scraping   adapter.TabelogInterface
		presenter  ScrapingSamplePresenter
		ctxTimeout time.Duration
	}
)

func NewScrapingSampleInteractor(
	scraping adapter.TabelogInterface,
	presenter ScrapingSamplePresenter,
	ctxTimeout time.Duration,
) ScrapingSampleUseCase {
	return &ScrapingSampleInteractor{
		scraping:   scraping,
		presenter:  presenter,
		ctxTimeout: ctxTimeout,
	}
}

func (i *ScrapingSampleInteractor) Execute(
	ctx context.Context,
	input ScrapingSampleInput,
) (ScrapingSampleOutput, error) {
	ctx, cancel := context.WithTimeout(ctx, i.ctxTimeout)
	defer cancel()

	// スクレイピングはinfra層に関連する処理だと思うので
	// より具体的な内容はdomain-serviceで実装？
	err := i.scraping.GetSample(ctx, input.Url)
	if err != nil {
		return ScrapingSampleOutput{}, err
	}

	return i.presenter.Output(), nil
}
