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

type HowMuchIOweRequest struct {
	MobileNumber string `json:"mobileNumber"`
}

type HowMuchIOweResponse struct {
	Transactions []splitmodels.Debt `json:"transactions"`
	Err          error              `json:"-"`
}

func (h *HowMuchIOweResponse) Failed() error { return h.Err }

type HowMuchOthersOweToMeRequest struct {
	MobileNumber string `json:"mobileNumber"`
}

type HowMuchOthersOweToMeResponse struct {
	Transactions []splitmodels.Transaction `json:"transactions"`
	Err          error                     `json:"-"`
}

func (h *HowMuchOthersOweToMeResponse) Failed() error { return h.Err }

type ChangePaymentStatusRequest struct {
	MobileNumber          string `json:"mobileNumber"`
	TransactionIdentifier string `json:"transactionIdentifier"`
}

type ChangePaymentStatusResponse struct {
	IsUpdated bool  `json:"isUpdated"`
	Err       error `json:"-"`
}

func (h *ChangePaymentStatusResponse) Failed() error { return h.Err }
