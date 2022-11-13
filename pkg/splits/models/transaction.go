package splitmodels

import (
	"github.com/google/uuid"
)

type Transaction struct {
	Id                   interface{} `json:"_id,omitempty" bson:"_id,omitempty"`
	Identifier           string      `json:"identifier"`
	Place                string      `json:"place"`
	TotalAmount          float64     `json:"totalAmount"`
	Date                 string      `json:"date"`
	SpentBy              SpentBy     `json:"spentBy"`
	NPeople              int         `json:"nPeople"`
	Split                []Split     `json:"split"`
	OverallPaymentStatus string      `json:"overallPaymentStatus"`
}

func (t *Transaction) InitFields() {
	t.Identifier = uuid.New().String()
}
