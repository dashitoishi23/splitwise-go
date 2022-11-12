package splitmodels

type Debt struct {
	TotalAmount   float64       `bson:"floatAmount"`
	Place         string        `bson:"place"`
	Date          string        `bson:"date"`
	SpentBy       SpentBy       `bson:"spentBy"`
	Npeople       int           `bson:"nPeople"`
	MyShare       float64       `bson:"myShare"`
	PaymentStatus PaymentStatus `bson:"paymentStatus"`
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
