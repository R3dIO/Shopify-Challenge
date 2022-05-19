package controllers

import "github.com/R3dIO/shopify-production_engineer/models"

type CreateItemInput struct {
	Name     string `json:"name" binding:"required"`
	City     string `json:"city" binding:"required,oneof=London HongKong Singapore"`
	Quantity int    `json:"quantity"`
}

type UpdateItemInput struct {
	Name     string `json:"name"`
	City     string `json:"city"`
	Quantity int    `json:"quantity"`
}

func UpdateHttpReqToDBReq(req UpdateItemInput) models.Item {
	return models.Item{
		Name: req.Name,
		City: req.City,
		Quantity: req.Quantity,
	}
}
