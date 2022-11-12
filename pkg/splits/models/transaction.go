package splitmodels

import (
	"time"

	"github.com/google/uuid"
)

type Transaction struct {
	Identifier           string        `bson:"identifier"`
	Place                string        `bson:"place"`
	TotalAmount          int           `bson:"totalAmount"`
	Date                 time.Time     `bson:"date"`
	SpentBy              SpentBy       `bson:"spentBy"`
	NPeople              int           `bson:"nPeople"`
	Split                []Split       `bson:"split"`
	OverallPaymentStatus PaymentStatus `bson:"overallPaymentStatus"`
}

func (t *Transaction) InitFields() {
	t.Identifier = uuid.New().String()
}
