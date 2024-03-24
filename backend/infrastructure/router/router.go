package router

import (
	"context"
	"fmt"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	action "github.com/yosuke7040/dine-out-discoveries/adapter/api"
	"github.com/yosuke7040/dine-out-discoveries/adapter/presenter"
	adapterScraping "github.com/yosuke7040/dine-out-discoveries/adapter/scraping"
	"github.com/yosuke7040/dine-out-discoveries/usecase"
)

type serverMuxEngine struct {
	router     *http.ServeMux
	port       Port
	ctxTimeout time.Duration
	scraping   adapterScraping.CollyScraping
	// scraping   adapterScraping.Scraping
}

func newServerMuxEngine(
	port Port,
	t time.Duration,
	scraping adapterScraping.CollyScraping,
) *serverMuxEngine {
	return &serverMuxEngine{
		router:     http.NewServeMux(),
		port:       port,
		ctxTimeout: t,
		scraping:   scraping,
	}
}

func (s *serverMuxEngine) Listen() {
	slog.Info("Server is running... ", "on port", s.port)

	s.setAppHandlers()

	server := &http.Server{
		Addr:              fmt.Sprintf(":%d", s.port),
		Handler:           s.router,
		ReadHeaderTimeout: 10 * time.Second,
		ReadTimeout:       10 * time.Minute,
		WriteTimeout:      10 * time.Minute,
	}

	go func() {
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			panic(err)
		}
	}()

	// Graceful shutdown
	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)
	<-stop

	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer func() {
		cancel()
	}()

	if err := server.Shutdown(ctx); err != nil {
		slog.Error("Server forced to shutdown: ", err)
		os.Exit(1)
	}

	slog.Info("Server shutdown properly")
}

func (s *serverMuxEngine) setAppHandlers() {
	s.router.HandleFunc("GET /", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello, World!!"))
	})
	s.router.HandleFunc("GET /healthcheck", s.healthcheck())

	// s.router.HandleFunc("GET /scraping/restaurant", s.buildScrapingRestaurantAction())
	s.router.HandleFunc("GET /scraping/sample", s.buildScrapingSampleAction())
}

func (s *serverMuxEngine) buildScrapingSampleAction() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		uc := usecase.NewScrapingSampleInteractor(
			// repository.NewRestaurantRepository(),
			// s.scraping,
			adapterScraping.NewTabelogScraping(s.scraping),
			presenter.NewScrapingSamplePresenter(),
			s.ctxTimeout,
		)
		act := action.NewScrapingSampleAction(uc)

		act.Execute(w, r)
	}
}

func (s *serverMuxEngine) healthcheck() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		action.HealthCheck(w, r)
	}
}
