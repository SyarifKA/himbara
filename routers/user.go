package routers

import (
	"github.com/SyarifKA/himbara/controllers"
	"github.com/gin-gonic/gin"
)

func UserRouters(r *gin.RouterGroup) {
	r.POST("", controllers.CheckUser)
	r.POST("/purchase", controllers.PurchaseOrder)
}
