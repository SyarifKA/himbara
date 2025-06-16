package controllers

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"

	"github.com/SyarifKA/himbara/dtos"
	"github.com/SyarifKA/himbara/lib"
	"github.com/SyarifKA/himbara/logs/logger"
	"github.com/SyarifKA/himbara/service"
	"github.com/gin-gonic/gin"
)

func CheckUser(ctx *gin.Context) {
	form := dtos.User{}

	ctx.ShouldBind(&form)

	jsonData, err := json.Marshal(form)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "failed to marshal data"})
		return
	}

	// Kirim ke endpoint lain
	resp, err := http.Post("http://localhost:8080/users", "application/json", bytes.NewBuffer(jsonData))
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "failed to send request"})
		return
	}
	defer resp.Body.Close()

	// // Baca respons dari endpoint tujuan
	body, _ := io.ReadAll(resp.Body)

	// isEligible := string(body)

	// // logger.Info("masuk")

	// Parse JSON response
	var isEligible map[string]interface{}
	if err := json.Unmarshal(body, &isEligible); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "failed to parse response"})
		return
	}

	success := isEligible["success"].(bool)

	// Parse JSON response
	var result map[string]interface{}
	if err := json.Unmarshal(body, &result); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "failed to parse response"})
		return
	}

	data := result["results"].([]interface{})
	if !success {
		lib.HandlerBadReq(ctx, "Maaf permintaan anda tidak dapat diproses")
		return
	} else {
		// result := service.CheckoutOrder(form)
		logger.Info("List paket tersedia")
		lib.HandlerOK(ctx, "Transaksi dapat dilanjutkan", data, nil)
	}
}

func PurchaseOrder(ctx *gin.Context) {
	form := dtos.PurchaseOrder{}
	ctx.ShouldBind(&form)

	result, err := service.PurchaseOrder(form)
	if err != nil {
		lib.HandlerBadReq(ctx, "Failed to purchase order")
		return
	}
	// fmt.Println(result)
	lib.HandlerOK(ctx, "Transaction Success", result, nil)
}
