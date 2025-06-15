package routers

import "github.com/gin-gonic/gin"

func RoutersCombine(r *gin.Engine) {
	UserRouters(r.Group("users"))
	Midtrans(r.Group("midtrans"))
}
