package splitmodels

type Debt struct {
	TotalAmount   float64 `json:"floatAmount"`
	Place         string  `json:"place"`
	Date          string  `json:"date"`
	SpentBy       SpentBy `json:"spentBy"`
	Npeople       int     `json:"nPeople"`
	MyShare       float64 `json:"myShare"`
	PaymentStatus string  `json:"paymentStatus"`
}

type PaymentStatus int64

const (
	Paid PaymentStatus = iota
	Pending
)

func (s PaymentStatus) String() string {
	switch s {
	case Paid:
		return "Paid"
	case Pending:
		return "Pending"
	}
	return "unknown"
}
