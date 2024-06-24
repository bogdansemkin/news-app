package api

import (
	"github.com/rs/zerolog"
	"net/http"
	"news-app/pkg/model"
	"strconv"
)

func (a *API) PostGetAllHandler(w http.ResponseWriter, r *http.Request) {
	limit := r.URL.Query().Get("limit")
	offset := r.URL.Query().Get("offset")

	if limit == "" {
		limit = "0"
	}
	limitArg, err := strconv.Atoi(limit)
	if err != nil {
		zerolog.Ctx(r.Context()).Error().Err(err).Msg("[PostGetAllHandler] converting limit")

		err = JSON(w, []any{})
		if err != nil {
			zerolog.Ctx(r.Context()).Error().Err(err).Msg("[PostGetAllHandler] all posts")
		}

		return
	}
	if offset == "" {
		offset = "0"
	}
	offsetArg, err := strconv.Atoi(offset)
	if err != nil {
		zerolog.Ctx(r.Context()).Error().Err(err).Msg("[PostGetAllHandler] converting offset")

		err = JSON(w, []any{})
		if err != nil {
			zerolog.Ctx(r.Context()).Error().Err(err).Msg("[PostGetAllHandler] all posts")
		}

		return
	}
	posts, err := a.service.GetAll(r.Context(), model.PaginationArgs{
		Limit:  limitArg,
		Offset: offsetArg,
	})

	if err != nil {
		zerolog.Ctx(r.Context()).Error().Err(err).Msg("[PostGetAllHandler] all posts")

		err = JSON(w, []any{})
		if err != nil {
			zerolog.Ctx(r.Context()).Error().Err(err).Msg("[PostGetAllHandler] all posts")
		}

		return
	}

	err = JSON(w, posts)
	if err != nil {
		zerolog.Ctx(r.Context()).Error().Err(err).Msg("[PostGetAllHandler] all posts")
	}
}
