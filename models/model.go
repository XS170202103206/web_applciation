package models

import (
	"gorm.io/gorm"
)

type DemoOrder struct{
	gorm.Model
	OrderNo  string `json:"order_no"`
	UserName string `json:"username"`
	Amount   float64 `json:"amount"`
	Status  string `json:"status"`
	FileUrl string `json:"file_url"`
}

type DemoOrders struct {
	Models []*DemoOrder `json:"models"`
	//ms.Models []models.DemoOrder
}
