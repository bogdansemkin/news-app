package api

import (
	"context"
	"errors"
	"github.com/go-chi/chi/v5"
	"log"
	"log/slog"
	"net/http"
	"news-app/internal/config"
	"news-app/internal/service"
	"os"
	"os/signal"
	"syscall"
	"time"
)

type API struct {
	cfg     *config.Config
	router  *chi.Mux
	srv     *http.Server
	service *service.Service
}

func New(config *config.Config, service *service.Service) *API {
	r := chi.NewRouter()

	return &API{
		router:  r,
		service: service,
		cfg:     config,
	}
}

func (a *API) Serve() {
	srv := &http.Server{
		WriteTimeout: 10 * time.Second,
		ReadTimeout:  10 * time.Second,
		IdleTimeout:  30 * time.Second,
		Addr:         a.cfg.HTTP.Host + a.cfg.HTTP.Port,
		Handler:      a.router,
	}
	a.srv = srv

	wait := a.GracefulShutdown()

	err := srv.ListenAndServe()
	if err != nil && !errors.Is(err, http.ErrServerClosed) {
		slog.Error("server shutdown cause of err", slog.String("error", err.Error()))
	}
	<-wait
}

func (a *API) GracefulShutdown() <-chan struct{} {
	var wait = make(chan struct{}, 1)
	var sigCh = make(chan os.Signal, 1)
	signal.Notify(sigCh, os.Interrupt, syscall.SIGTERM)
	go func() {
		s := <-sigCh
		defer func() {
			wait <- struct{}{}
		}()

		ctx, cancel := context.WithTimeout(context.Background(), time.Second*30)
		defer cancel()

		go func() {
			<-ctx.Done()
			if errors.Is(ctx.Err(), context.DeadlineExceeded) {
				log.Println("graceful shutdown timed out.. forcing exit.")
				os.Exit(-1)
			}
		}()

		err := a.srv.Shutdown(ctx)
		if err != nil {
			log.Println("shutdown with err cause of os signal")
			return
		} else {
			log.Println("shutdown cause of os signal")
			return
		}

		s.Signal()
	}()

	return wait
}
