package controllers

import (
	"fmt"
	"net/http"

	"github.com/SyarifKA/himbara/lib"
	"github.com/SyarifKA/himbara/models"
	"github.com/gin-gonic/gin"
)

func MidtransNotification(c *gin.Context) {
	var notifPayload map[string]interface{}
	if err := c.BindJSON(&notifPayload); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid payload"})
		return
	}

	fmt.Println(notifPayload)

	Id := notifPayload["order_id"].(string)
	transactionStatus := notifPayload["transaction_status"].(string)
	paymentType := notifPayload["payment_type"].(string)

	fmt.Println(transactionStatus)

	// Update ke database
	db := lib.ConnectDB()
	var po models.PurchaseOrder
	if err := db.First(&po, "id = ?", Id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Order not found"})
		return
	}

	po.Status = transactionStatus
	po.PaymentChanel = paymentType
	db.Save(&po)

	c.JSON(http.StatusOK, gin.H{"message": "Status updated"})
}
