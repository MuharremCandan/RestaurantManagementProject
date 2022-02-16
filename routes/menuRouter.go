package routes

import (
	controller "golangRestaurantManagement/controllers"

	"github.com/gin-gonic/gin"
)

func MenuRoutes(incomingRoutes *gin.Engine) {

	incomingRoutes.GET("/menues", controller.GetMenus())
	incomingRoutes.GET("menues/:menu_id", controller.GetMenu())
	incomingRoutes.POST("/menues", controller.CreateMenu())
	incomingRoutes.PATCH("/menues/:menu_id", controller.UpdateMenu())
}
