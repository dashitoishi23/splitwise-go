package util

import (
	"context"
	"encoding/json"
	"io"
	"net/http"

	"github.com/go-kit/kit/endpoint"
)

func EncodeHTTPGenericResponse(ctx context.Context, w http.ResponseWriter, response interface{}) error {
	if e, ok := response.(endpoint.Failer); ok && e.Failed() != nil {
		ErrorEncoder(ctx, e.Failed(), w)

		return nil
	}

	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Content-Type", "application/json")
	return json.NewEncoder(w).Encode(response)
}

func DecodeHTTPGenericRequest[T any](ctx context.Context, r *http.Request) (interface{}, error) {

	var req T

	defer r.Body.Close()

	decodedReq := json.NewDecoder(r.Body)

	decodedReq.DisallowUnknownFields()

	err := decodedReq.Decode(&req)

	if err == io.EOF {
		return req, nil
	}

	if err != nil {
		return req, err
	}

	return req, nil

}
