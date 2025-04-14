package domain

import "time"

type TandaUsuario struct {
	ID           int       `json:"id"`
	TandaID      int       `json:"tandaId"`
	MemberID     int       `json:"memberId"`
	NumberTicket int       `json:"numberTicket"`
	DatePay      time.Time `json:"datePay"`
	Status       string    `json:"status"`
}
