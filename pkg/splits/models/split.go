package splitmodels

type Split struct {
	Mobile        string        `json:"mobile"`
	Name          string        `json:"string"`
	ShareAmount   string        `json:"shareAmount"`
	PaymentStatus PaymentStatus `json:"paymentStatus"`
}
