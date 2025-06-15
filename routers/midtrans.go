package routers

import (
	"github.com/SyarifKA/himbara/controllers"
	"github.com/SyarifKA/himbara/middlewares"
	"github.com/gin-gonic/gin"
)

func Midtrans(r *gin.RouterGroup) {
	r.POST("/notification", middlewares.LogWithMessage("notif midtrans", controllers.MidtransNotification))
}
