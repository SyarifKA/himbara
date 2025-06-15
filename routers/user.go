package routers

import (
	"github.com/SyarifKA/himbara/controllers"
	"github.com/SyarifKA/himbara/middlewares"
	"github.com/gin-gonic/gin"
)

func UserRouters(r *gin.RouterGroup) {
	r.POST("", middlewares.LogWithMessage("Check user success", controllers.CheckUser))
	r.POST("/purchase", middlewares.LogWithMessage("Purchase order success", controllers.PurchaseOrder))
}
