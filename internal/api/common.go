package api

import (
	"encoding/json"
	"net/http"

	"github.com/pkg/errors"
)

func JSON(w http.ResponseWriter, n any) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	b, err := json.Marshal(n)
	if err != nil {
		return errors.Wrap(err, "marshal failed")
	}

	_, err = w.Write(b)
	if err != nil {
		return errors.Wrap(err, "write body failed")
	}

	return nil
}

func JSONCode(w http.ResponseWriter, n any, statusCode int) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)

	b, err := json.Marshal(n)
	if err != nil {
		return errors.Wrap(err, "marshal failed")
	}

	_, err = w.Write(b)
	if err != nil {
		return errors.Wrap(err, "write body failed")
	}

	return nil
}
