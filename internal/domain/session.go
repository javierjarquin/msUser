package domain

import "time"

type Session struct {
	ID           int       `json:"id"`
	UserID       int       `json:"userId"`
	CreationDate time.Time `json:"creationDate"`
	IPAddres     string    `json:"ipAddres"`
	Comments     string    `json:"comments"`
}
