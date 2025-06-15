package models

import "time"

type ListPackage struct {
	PaketID   string `json:"paketId"`
	PaketName string `json:"paketName"`
	Harga     int64  `json:"harga"`
}

type Purchase struct {
	PaketName string `json:"paketName"`
}
type PurchaseOrder struct {
	Id            string
	PhoneNumber   string
	PaymentChanel string
	ProductId     string
	ProductName   string
	Amount        int64
	Status        string
	CreatedAt     time.Time
}

func (PurchaseOrder) TableName() string {
	return "purchases"
}
