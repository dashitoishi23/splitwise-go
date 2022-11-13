package util

import (
	"context"
	"encoding/json"
	"net/http"
)

func ErrorEncoder(_ context.Context, err error, w http.ResponseWriter) {
	w.WriteHeader(getHeader(err.Error()))

	json.NewEncoder(w).Encode(errorWrapper{Error: err.Error()})
}

func getHeader(errMsg string) int {
	switch errMsg {
	case DEBT_NOT_FOUND:
		return http.StatusNotFound
	case TRANSACTION_NOT_FOUND:
		return http.StatusNotFound
	default:
		return http.StatusInternalServerError
	}
}

type errorWrapper struct {
	Error string `json:"error"`
}
