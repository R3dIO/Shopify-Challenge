package controllers

import "github.com/R3dIO/shopify-production_engineer/models"
var cities = "London Vancouver Delhi Tokyo Houston"
type CreateItemInput struct {
	Name     string `json:"name" binding:"required"`
	City     string `json:"city" binding:"required,oneof=London Vancouver Delhi Tokyo Houston"`
	Quantity int    `json:"quantity"`
}

type UpdateItemInput struct {
	Name     string `json:"name"`
	City     string `json:"city" binding:"oneof=London Vancouver Delhi Tokyo Houston"`
	Quantity int    `json:"quantity"`
}

func UpdateHttpReqToDBReq(req UpdateItemInput) models.Item {
	return models.Item{
		Name: req.Name,
		City: req.City,
		Quantity: req.Quantity,
	}
}

type Coordinates struct {
	Name     	string  `json:"name"`
	Lat			float64  `json:"lat"`
	Lon			float64  `json:"lon"`
	Country		string  `json:"country"`
}