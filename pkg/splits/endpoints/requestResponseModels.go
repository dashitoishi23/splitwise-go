package splitendpoints

import (
	splitmodels "github.com/dashitoishi23/splitwise-go/pkg/splits/models"
)

type SaveTheTransactionRequest struct {
	Transaction splitmodels.Transaction `json:"transaction"`
}

type SaveTheTransactionResponse struct {
	Err error `json:"-"`
}

func (s *SaveTheTransactionResponse) Failed() error { return s.Err }
