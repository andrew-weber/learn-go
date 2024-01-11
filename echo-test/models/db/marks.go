package models

type Mark struct {
	ID     int    `json:"id"`
	Date   string `json:"date"`
	Active bool   `json:"active"`
}
