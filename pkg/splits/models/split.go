package splitmodels

type Split struct {
	Mobile        string        `json:"mobile"`
	Name          string        `json:"name"`
	ShareAmount   float64       `json:"shareAmount"`
	PaymentStatus PaymentStatus `json:"paymentStatus"`
}
