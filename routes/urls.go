package routes

import (
	"century/controllers"
	"github.com/gin-gonic/gin"
)

func AddPublicAPIs(r *gin.RouterGroup, controller *controllers.BankController) {
	r.POST("/transfer", controller.TransferHandler)
	r.GET("/balance", controller.BalanceHandler)
}

func AddMainAPIs(r *gin.Engine, controller *controllers.BankController) {
	AddPublicAPIs(r.Group("/api/v1"), controller)
}
