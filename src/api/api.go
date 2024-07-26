package api

import (
	"safe-plate/src/api/controllers"
	"safe-plate/src/middleware"

	"github.com/gin-gonic/gin"
)

func InitServer() {
	r := gin.Default()

	r.POST("/signup", controllers.Signup)
	r.POST("/login", controllers.Login)
	r.GET("/validate", middleware.RequireAuth, controllers.Validate)

	r.POST("/products", controllers.ProductCreate)
	r.PUT("/products/:id", controllers.ProductUpdate)
	r.GET("/products", controllers.ProductsIndex)
	r.GET("/products/:id", controllers.ProductShow)
	r.DELETE("/products/:id", controllers.ProductDelete)

	r.Run()
}
