package controllers

import (
	"fmt"
	"net/http"
	"os"
	"github.com/R3dIO/shopify-production_engineer/models"
	"github.com/gin-gonic/gin"
)

// GET /item
// Get all items
func ListItems(c *gin.Context) {
  var items []models.Item
  models.DB.Find(&items)

  c.JSON(http.StatusOK, gin.H{"data": items})
}


func CreateItem(c *gin.Context) {
	// Validate input
	var input CreateItemInput
	if err := c.ShouldBindJSON(&input); err != nil {
	  c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	  return
	}

	key := os.Getenv("KEY")  
	// Create Item
	item := models.Item{Name: input.Name, City: input.City, Quantity: input.Quantity}
	coordinates := []Coordinates{}
	err := GetJsonRespFromUrl("http://api.openweathermap.org/geo/1.0/direct?q=London&limit=5&appid=" + key, &coordinates, false)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	} else {
		fmt.Printf("resp %+v\n", coordinates)
	}

	lat := fmt.Sprintf("%f", coordinates[0].Lat) 
	lon := fmt.Sprintf("%f", coordinates[0].Lon) 

	body, err := GetStringRespFromUrl("https://api.openweathermap.org/data/2.5/weather?lat=" + lat + "&lon= "+ lon + "&appid=" + key, false)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	item.Description = body
	err = models.CreateItem(&item)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": item})
}


func FindItemById(c *gin.Context) {  // Get model if exist
	var item models.Item
	Id := c.Param("id")
	err := models.GetOneItem(Id, &item)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": item})
}

// PATCH /items/:id
// Update a book
func UpdateItem(c *gin.Context) {
	// Get model if exist
	var item models.Item
	Id := c.Param("id")

	err := models.GetOneItem(Id, &item)
	if err != nil {
	  c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
	  return
	}
  
	// Validate input
	var inputRaw UpdateItemInput
	if err := c.ShouldBindJSON(&inputRaw); err != nil {
	  c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	  return
	}
  
	input := UpdateHttpReqToDBReq(inputRaw)
	err = models.UpdateItem(&item, &input)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": item})
}

func DeleteItem(c *gin.Context) {
	// Get model if exist
	var item models.Item
	Id := c.Param("id")
	
	err := models.GetOneItem(Id, &item)
	if err != nil {
	  c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
	  return
	}
  
	err = models.DeleteItem(&item)
	if  err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"data": false})
		return
	}
  
	c.JSON(http.StatusOK, gin.H{"data": true})
  }