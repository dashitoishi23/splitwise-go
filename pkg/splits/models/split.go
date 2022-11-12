package splitmodels

type Split struct {
	Mobile        string        `bson:"mobile"`
	Name          string        `bson:"string"`
	ShareAmount   string        `bson:"shareAmount"`
	PaymentStatus PaymentStatus `bson:"paymentStatus"`
}
