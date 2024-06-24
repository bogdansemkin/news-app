package api

import (
	"github.com/go-chi/chi/v5"
	"github.com/rs/zerolog"
	"net/http"
	"news-app/pkg/model"
	"strconv"
)

func (a *API) PostGetByIdHandler(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	limit := r.URL.Query().Get("limit")
	offset := r.URL.Query().Get("offset")

	idArg, err := strconv.Atoi(id)
	if err != nil {
		zerolog.Ctx(r.Context()).Error().Err(err).Msg("[PostGetByIdHandler] converting id")

		err = JSON(w, []any{})
		if err != nil {
			zerolog.Ctx(r.Context()).Error().Err(err).Msg("[PostGetByIdHandler] converting id")
		}

		return
	}

	if limit == "" {
		limit = "0"
	}
	limitArg, err := strconv.Atoi(limit)
	if err != nil {
		zerolog.Ctx(r.Context()).Error().Err(err).Msg("[PostGetByIdHandler] converting limit")

		err = JSON(w, []any{})
		if err != nil {
			zerolog.Ctx(r.Context()).Error().Err(err).Msg("[PostGetByIdHandler] converting limit")
		}

		return
	}
	if offset == "" {
		offset = "0"
	}
	offsetArg, err := strconv.Atoi(offset)
	if err != nil {
		zerolog.Ctx(r.Context()).Error().Err(err).Msg("[PostGetByIdHandler] converting offset")

		err = JSON(w, []any{})
		if err != nil {
			zerolog.Ctx(r.Context()).Error().Err(err).Msg("[PostGetByIdHandler] converting offset")
		}

		return
	}

	post, err := a.service.GetById(r.Context(), model.PaginationByID{
		ID: idArg,
		PaginationArgs: model.PaginationArgs{
			Limit:  limitArg,
			Offset: offsetArg,
		},
	})

	if err != nil {
		zerolog.Ctx(r.Context()).Error().Err(err).Msg("[PostGetByIdHandler] all posts")

		err = JSON(w, struct{}{})
		if err != nil {
			zerolog.Ctx(r.Context()).Error().Err(err).Msg("[PostGetByIdHandler] all posts")
		}

		return
	}

	err = JSON(w, post)
	if err != nil {
		zerolog.Ctx(r.Context()).Error().Err(err).Msg("[PostGetByIdHandler] all posts")
	}
}
