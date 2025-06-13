package models

type ListPackage struct {
	PaketID   string `json:"paketId"`
	PaketName string `json:"paketName"`
	Harga     int64  `json:"harga"`
}
