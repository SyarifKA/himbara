package service

import (
	"bytes"
	"encoding/json"
	"net/http"

	"github.com/SyarifKA/himbara/dtos"
	"github.com/SyarifKA/himbara/lib"
	"github.com/SyarifKA/himbara/models"
	"github.com/google/uuid"
)

func CheckoutOrder(form dtos.User) []models.ListPackage {
	var result []models.ListPackage

	jsonData, err := json.Marshal(form)
	if err != nil {
		return result // return kosong jika gagal marshal
	}

	resp, err := http.Post("http://localhost:8080/receiver", "application/json", bytes.NewBuffer(jsonData))
	if err != nil {
		return result // return kosong jika gagal kirim
	}
	defer resp.Body.Close()

	// decode body response ke struct result
	json.NewDecoder(resp.Body).Decode(&result)
	return result
}

func PurchaseOrder(order dtos.PurchaseOrder) (*models.PurchaseOrder, error) {
	db := lib.ConnectDB()
	status := "waiting"

	// Simpan ke database lokal
	po := models.PurchaseOrder{
		Id:            uuid.New().String(),
		PhoneNumber:   order.PhoneNumber,
		PaymentChanel: order.PaymentChanel,
		ProductId:     order.ProductId,
		ProductName:   order.ProductName,
		Amount:        order.Amount,
		Status:        status,
	}

	if err := db.Create(&po).Error; err != nil {
		return nil, err
	}

	// Siapkan data untuk dikirim ke endpoint
	// payload := dtos.CheckoutPayload{
	// 	ProductId:     order.ProductId,
	// 	ProductName:   order.ProductName,
	// 	PaymentChanel: order.PaymentChanel,
	// 	Amount:        order.Amount,
	// }

	// jsonData, err := json.Marshal(payload)
	// if err != nil {
	// 	return &po, err
	// }

	// Kirim ke endpoint lain
	// resp, err := http.Post("http://localhost:8080/receiver", "application/json", bytes.NewBuffer(jsonData))
	// if err != nil {
	// 	return &po, err
	// }
	// defer resp.Body.Close()

	// Bisa decode respons kalau perlu, misalnya:
	// var endpointResponse dtos.SomeResponse
	// _ = json.NewDecoder(resp.Body).Decode(&endpointResponse)

	return &po, nil
}
