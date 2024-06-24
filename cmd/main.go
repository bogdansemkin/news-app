package main

import (
	"github.com/rs/zerolog/log"
	"news-app/internal/api"
	"news-app/internal/config"
	"news-app/internal/repository/db/postgres"
	"news-app/internal/repository/post"
	service2 "news-app/internal/service"
)

func main() {
	cfg, err := config.LoadConfig()
	if err != nil {
		panic(err)
	}
	db, err := postgres.NewPostgresDB(cfg.DB.PostgresConn)
	if err != nil {
		log.Panic().Err(err).Send()
	}

	repos := post.NewRepository(db)
	service := service2.NewService(repos)
	apiServer, err := api.New(
		cfg,
		service,
	).
		InitRoutes()
	if err != nil {
		log.Panic().Err(err).Send()
	}

	log.Info().Msg("Listening on " + cfg.HTTP.Port)

	apiServer.Serve()
}
