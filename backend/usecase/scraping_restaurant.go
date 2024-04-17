package usecase

import (
	"context"
	"log/slog"
	"time"

	adapter "github.com/yosuke7040/dine-out-discoveries/adapter/scraping"
)

type (
	ScrapingRestaurantUseCase interface {
		Execute(context.Context, ScrapingRestaurantInput) (ScrapingRestaurantOutput, error)
	}

	ScrapingRestaurantInput struct {
		Url string `json:"url"`
	}

	ScrapingRestaurantPresenter interface {
		Output() ScrapingRestaurantOutput
	}

	ScrapingRestaurantOutput struct {
		Result string `json:"id"`
	}

	ScrapingRestaurantInteractor struct {
		scraping   adapter.TabelogInterface
		presenter  ScrapingRestaurantPresenter
		ctxTimeout time.Duration
	}
)

func NewScrapingRestaurantInteractor(
	scraping adapter.TabelogInterface,
	presenter ScrapingRestaurantPresenter,
	ctxTimeout time.Duration,
) ScrapingRestaurantUseCase {
	return &ScrapingRestaurantInteractor{
		scraping:   scraping,
		presenter:  presenter,
		ctxTimeout: ctxTimeout,
	}
}

func (i *ScrapingRestaurantInteractor) Execute(
	ctx context.Context,
	input ScrapingRestaurantInput,
) (ScrapingRestaurantOutput, error) {
	ctx, cancel := context.WithTimeout(ctx, i.ctxTimeout)
	defer cancel()

	err := i.scraping.GetRestaurantTopPage(ctx, input.Url)
	if err != nil {
		return ScrapingRestaurantOutput{}, err
	}

	urlLists, err := i.scraping.GetReviewUrlLists(ctx, input.Url)
	if err != nil {
		return ScrapingRestaurantOutput{}, err
	}
	slog.Info("ScrapingRestaurantInteractor", "urlLists", urlLists)

	// ? クリーンアーキテクチャってループしていいか？
	for _, url := range urlLists {
		slog.Info("ScrapingRestaurantInteractor", "url", url)
		err = i.scraping.GetReview(ctx, url)
		if err != nil {
			return ScrapingRestaurantOutput{}, err
		}
	}

	return i.presenter.Output(), nil
}
