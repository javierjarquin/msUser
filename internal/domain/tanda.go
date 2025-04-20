package domain

import "time"

type Tanda struct {
	ID             int       `json:"id"`
	Alias          string    `json:"alias"`
	PoolAmount     float64   `json:"poolAmount"`
	Period         string    `json:"period"`
	Members        int       `json:"members"`
	StartDate      time.Time `json:"startDate"`
	EndDate        time.Time `json:"endDate"`
	TotalEndPool   float64   `json:"totalEndPool"`
	CreationUserId int       `json:"creationUserId"`
	UserCreation   string    `json:"userCreation"`
}
