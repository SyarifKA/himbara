package service

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"os"

	"github.com/SyarifKA/himbara/dtos"
	"github.com/SyarifKA/himbara/lib"
	"github.com/SyarifKA/himbara/models"
	"github.com/google/uuid"
	"github.com/veritrans/go-midtrans"
)

func CheckoutOrder(form dtos.User) []models.ListPackage {
	var result []models.ListPackage

	jsonData, err := json.Marshal(form)
	if err != nil {
		return result // return kosong jika gagal marshal
	}

	resp, err := http.Post("http://localhost:8080/users", "application/json", bytes.NewBuffer(jsonData))
	if err != nil {
		return result // return kosong jika gagal kirim
	}
	defer resp.Body.Close()

	// decode body response ke struct result
	json.NewDecoder(resp.Body).Decode(&result)
	return result
}

// func PurchaseOrder(order dtos.PurchaseOrder) (*models.PurchaseOrder, error) {
// 	db := lib.ConnectDB()
// 	status := "waiting"

// 	// Simpan ke database lokal
// 	po := models.PurchaseOrder{
// 		Id:            uuid.New().String(),
// 		PhoneNumber:   order.PhoneNumber,
// 		PaymentChanel: order.PaymentChanel,
// 		ProductId:     order.ProductId,
// 		ProductName:   order.ProductName,
// 		Amount:        order.Amount,
// 		Status:        status,
// 	}

// 	if err := db.Create(&po).Error; err != nil {
// 		return nil, err
// 	}

// 	// Siapkan data untuk dikirim ke endpoint
// 	// payload := dtos.CheckoutPayload{
// 	// 	ProductId:     order.ProductId,
// 	// 	ProductName:   order.ProductName,
// 	// 	PaymentChanel: order.PaymentChanel,
// 	// 	Amount:        order.Amount,
// 	// }

// 	// jsonData, err := json.Marshal(payload)
// 	// if err != nil {
// 	// 	return &po, err
// 	// }

// 	// Kirim ke endpoint lain
// 	// resp, err := http.Post("http://localhost:8080/receiver", "application/json", bytes.NewBuffer(jsonData))
// 	// if err != nil {
// 	// 	return &po, err
// 	// }
// 	// defer resp.Body.Close()

// 	// Bisa decode respons kalau perlu, misalnya:
// 	// var endpointResponse dtos.SomeResponse
// 	// _ = json.NewDecoder(resp.Body).Decode(&endpointResponse)

// 	return &po, nil
// }

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

	// 🔻 Integrasi Midtrans Snap di sini
	client := midtrans.NewClient()
	client.ServerKey = os.Getenv("MIDTRANS_SERVER_KEY")
	client.ClientKey = os.Getenv("MIDTRANS_CLIENT_KEY")
	client.APIEnvType = midtrans.Sandbox

	snap := midtrans.SnapGateway{Client: client}

	customer := midtrans.CustDetail{
		FName: order.PhoneNumber,
		Email: fmt.Sprintf("%s@example.com", order.PhoneNumber), // dummy email
		Phone: order.PhoneNumber,
	}

	items := []midtrans.ItemDetail{
		{
			ID:    order.ProductId,
			Name:  order.ProductName,
			Price: order.Amount,
			Qty:   1,
		},
	}

	req := &midtrans.SnapReq{
		TransactionDetails: midtrans.TransactionDetails{
			OrderID:  po.Id,
			GrossAmt: order.Amount,
		},
		CustomerDetail: &customer,
		Items:          &items,
	}

	snapResp, err := snap.GetToken(req)
	if err != nil {
		fmt.Println("Midtrans Snap error:", err.Error())
		return nil, err
	}

	// Dapatkan status transaksi dari Midtrans
	// core := midtrans.CoreGateway{Client: client}
	// txStatus, err := core.Status(po.Id)
	// if err != nil {
	// 	fmt.Println("Midtrans Status Check error:", err.Error())
	// 	return nil, err
	// }

	// po.Status = txStatus.TransactionStatus
	// po.PaymentChanel = txStatus.PaymentType

	// Simpan snap token & URL ke DB (opsional)
	po.SnapToken = snapResp.Token
	po.SnapRedirectUrl = snapResp.RedirectURL
	db.Save(&po)

	return &po, nil
}

func UpdateOrderStatus(Id, status, paymentType string) error {
	db := lib.ConnectDB()
	var po models.PurchaseOrder
	if err := db.First(&po, "id = ?", Id).Error; err != nil {
		return err
	}

	po.Status = status
	po.PaymentChanel = paymentType
	return db.Save(&po).Error
}
