package routes

import (
	controller "golangRestaurantManagement/controllers"

	"github.com/gin-gonic/gin"
)

func OrderItemroutes(incomingRoutes *gin.Engine) {

	incomingRoutes.GET("/orerItems", controller.GetOrderItems())
	incomingRoutes.GET("orerItems/:orderItem_id", controller.GetOrderItem())
	incomingRoutes.GET("/orderItems-order/:order_id", controller.GetOrderItemsByOrder())
	incomingRoutes.POST("/orerItems", controller.CreateOrderItem())
	incomingRoutes.PATCH("/orerItems/:orderItem_id", controller.UpdateOrderItem())
}
