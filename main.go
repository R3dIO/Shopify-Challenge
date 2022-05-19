package main

import (
	"net/http"

	"github.com/R3dIO/shopify-production_engineer/controllers"
	"github.com/R3dIO/shopify-production_engineer/models"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	models.ConnectDatabase()

	r.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"data": "Shopify production Intern Challenge"})    
	})

  	items := r.Group("/items"); {
		items.GET("/", controllers.ListItems)
		items.POST("/", controllers.CreateItem)
		items.GET("/:id", controllers.FindItemById)
		items.PATCH("/:id", controllers.UpdateItem)
		items.DELETE("/:id", controllers.DeleteItem)
	};

 	r.Run()
}