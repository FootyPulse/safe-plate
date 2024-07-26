package controllers

import (
	"encoding/json"
	"safe-plate/src/domain/models"
	"safe-plate/src/infra/persistence/database"

	"github.com/gin-gonic/gin"
)

func ProductCreate(c *gin.Context) {

	var body struct {
		Price      int
		Ingredient string
		Nutrition  json.RawMessage `json:"nutrition" gorm:"type:json"`
	}

	c.Bind(&body)

	product := models.Product{Price: body.Price, Ingredient: body.Ingredient, Nutrition: body.Nutrition}

	result := database.GetDb().Create(&product)

	if result.Error != nil {
		c.Status(400)
		return
	}

	c.JSON(200, gin.H{
		"product": product,
	})
}

func ProductsIndex(c *gin.Context) {

	var products []models.Product

	database.GetDb().Find(&products)

	c.JSON(200, gin.H{
		"products": products,
	})
}

func ProductShow(c *gin.Context) {
	id := c.Param("id")

	var product models.Product

	database.GetDb().First(&product, id)

	c.JSON(200, gin.H{
		"product": product,
	})
}

func ProductUpdate(c *gin.Context) {
	id := c.Param("id")

	var body struct {
		Price      int
		Ingredient string
		Nutrition  json.RawMessage `json:"nutrition" gorm:"type:json"`
	}

	c.Bind(&body)

	var product models.Product

	database.GetDb().First(&product, id)

	database.GetDb().Model(&product).Updates(models.Product{Price: body.Price, Ingredient: body.Ingredient, Nutrition: body.Nutrition})

	c.JSON(200, gin.H{
		"product": product,
	})
}

func ProductDelete(c *gin.Context) {
	id := c.Param("id")

	var product models.Product

	database.GetDb().Delete(&product, id)

	c.Status(200)
}
