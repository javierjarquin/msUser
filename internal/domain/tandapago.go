package domain

import "time"

type TandaPago struct {
	TandaUsuarioID int       `json:"tandaUsuarioId"`
	PeriodNumber   int       `json:"periodNumber"`
	PaymentDate    time.Time `json:"paymentDate"`
	Amount         float64   `json:"amount"`
	Status         string    `json:"status"`
}
