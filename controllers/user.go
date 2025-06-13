package controllers

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/SyarifKA/himbara/dtos"
	"github.com/SyarifKA/himbara/lib"
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
	resp, err := http.Post("http://localhost:8080/receiver", "application/json", bytes.NewBuffer(jsonData))
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "failed to send request"})
		return
	}
	defer resp.Body.Close()

	// Baca respons dari endpoint tujuan
	body, _ := io.ReadAll(resp.Body)
	fmt.Println(body)

	isEligible := string(body)

	if isEligible != "success" {
		lib.HandlerBadReq(ctx, "Maaf permintaan anda tidak dapat diproses")
		return
	} else {
		result := service.CheckoutOrder(form)
		lib.HandlerOK(ctx, "List paket tersedia", result, nil)
	}
}
