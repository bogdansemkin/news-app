package api

import (
	"github.com/go-chi/chi/v5"
	"github.com/rs/zerolog"
	"net/http"
	"news-app/pkg/model"
	"strconv"
)

func (a *API) PostDeleteHandler(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")

	idArg, err := strconv.Atoi(id)
	if err != nil {
		zerolog.Ctx(r.Context()).Error().Err(err).Msg("[PostDeleteHandler] converting id")

		err = JSONCode(w, map[string]string{"error": model.ErrPostInvalidId}, http.StatusBadRequest)
		if err != nil {
			zerolog.Ctx(r.Context()).Error().Err(err).Msg("[PostDeleteHandler] converting id")
		}

		return
	}

	err = a.service.Delete(r.Context(), idArg)
	if err != nil {
		zerolog.Ctx(r.Context()).Error().Err(err).Msg("[PostDeleteHandler] all posts")

		err = JSONCode(w, map[string]string{"error": err.Error()}, http.StatusInternalServerError)
		if err != nil {
			zerolog.Ctx(r.Context()).Error().Err(err).Msg("[PostCreateHandler] json failed")
		}

		return
	}

	err = JSONCode(w, []any{}, http.StatusOK)
	if err != nil {
		zerolog.Ctx(r.Context()).Error().Err(err).Msg("[PostCreateHandler] marshal response")
	}
}
