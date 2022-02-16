package main

import (
	"fmt"
	"golangRestaurantManagement/middleware"
	"golangRestaurantManagement/routes"
	"os"

	"github.com/gin-gonic/gin"
)

func init() {
	gin.SetMode(gin.ReleaseMode)
}

//var foodCollection *mongo.Collection = db.OpenCollection(db.Client, "food")

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = ":8080"
	}

	router := gin.New()
	gin.SetMode(gin.ReleaseMode)
	router.Use(gin.Logger())
	routes.UserRoutes(router)
	router.Use(middleware.Authentication())

	routes.FoodRoutes(router)
	routes.MenuRoutes(router)
	routes.TablaeRoutes(router)
	routes.OrderRoutes(router)
	routes.OrderItemroutes(router)
	routes.InvoiceRoutes(router)

	fmt.Println(port, ": connected")
	router.Run(port)

}
