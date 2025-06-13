package service

import (
	"bytes"
	"encoding/json"
	"net/http"

	"github.com/SyarifKA/himbara/dtos"
	"github.com/SyarifKA/himbara/models"
)

func CheckoutOrder(form dtos.User) models.ListPackage {
	var result models.ListPackage

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
