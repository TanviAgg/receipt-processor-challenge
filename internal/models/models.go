package models

import (
	"github.com/google/uuid"
)

type Points struct {
	Points int32 `json:"points"`
}

type Response struct {
	Id uuid.UUID `json:"id"`
}

type Item struct {
	ShortDescription string `json:"shortDescription"`
	Price            string `json:"price"`
}

type Purchase struct {
	Retailer     string `json:"retailer"`
	PurchaseDate string `json:"purchaseDate"`
	PurchaseTime string `json:"purchaseTime"`
	Items        []Item `json:"items"`
	Total        string `json:"total"`
}
