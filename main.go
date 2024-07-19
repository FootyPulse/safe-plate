package main

import (
	"safe-plate/controllers"
	"safe-plate/initializers"
	"safe-plate/middleware"

	"github.com/gin-gonic/gin"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectToDb()
	initializers.SyncDatabse()
}

func main() {
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
