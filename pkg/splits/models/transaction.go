package splitmodels

import "time"

type Transaction struct {
	Identifier  string    `json:"identifier"`
	Place       string    `json:"place"`
	TotalAmount int       `json:"totalAmount"`
	Date        time.Time `json:"date"`
	SpentBy     SpentBy   `json:"spentBy"`
	NPeople     int       `json:"nPeople"`
	Split       []Split   `json:"split"`
}
