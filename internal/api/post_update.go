package api

import (
	"encoding/json"
	"github.com/rs/zerolog"
	"net/http"
	"news-app/pkg/model"
)

func (a *API) PostUpdateHandler(w http.ResponseWriter, r *http.Request) {
	var postOnUpdate model.Post

	err := json.NewDecoder(r.Body).Decode(&postOnUpdate)
	if err != nil {
		zerolog.Ctx(r.Context()).Error().Err(err).Msg("[PostUpdateHandler]")

		err = JSONCode(w, map[string]string{"error": err.Error()}, http.StatusInternalServerError)
		if err != nil {
			zerolog.Ctx(r.Context()).Error().Err(err).Msg("[PostUpdateHandler] json failed")
		}
		return
	}

	result, err := a.service.Update(r.Context(), postOnUpdate)
	if err != nil {
		zerolog.Ctx(r.Context()).Error().Err(err).Msg("[PostUpdateHandler]")

		err = JSONCode(w, map[string]string{"error": err.Error()}, http.StatusInternalServerError)
		if err != nil {
			zerolog.Ctx(r.Context()).Error().Err(err).Msg("[PostCreateHandler] json failed")
		}

		return
	}

	err = JSONCode(w, result, http.StatusOK)
	if err != nil {
		zerolog.Ctx(r.Context()).Error().Err(err).Msg("[PostUpdateHandler] marshal response")
	}
}
