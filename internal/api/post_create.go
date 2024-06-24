package api

import (
	"encoding/json"
	"github.com/rs/zerolog"
	"net/http"
	"news-app/pkg/model"
)

func (a *API) PostCreateHandler(w http.ResponseWriter, r *http.Request) {
	var postOnCreate model.Post

	err := json.NewDecoder(r.Body).Decode(&postOnCreate)
	if err != nil {
		zerolog.Ctx(r.Context()).Error().Err(err).Msg("[PostCreateHandler]")

		err = JSONCode(w, map[string]string{"error": err.Error()}, http.StatusBadRequest)
		if err != nil {
			zerolog.Ctx(r.Context()).Error().Err(err).Msg("[PostCreateHandler] json failed")
		}

		return
	}

	err = a.service.Create(r.Context(), postOnCreate)
	if err != nil {
		zerolog.Ctx(r.Context()).Error().Err(err).Msg("[PostCreateHandler]")

		err = JSONCode(w, map[string]string{"error": err.Error()}, http.StatusInternalServerError)
		if err != nil {
			zerolog.Ctx(r.Context()).Error().Err(err).Msg("[PostCreateHandler] json failed")
		}

		return
	}

	err = JSONCode(w, struct{}{}, http.StatusOK)
	if err != nil {
		zerolog.Ctx(r.Context()).Error().Err(err).Msg("[PostCreateHandler] marshal response")
	}
}
