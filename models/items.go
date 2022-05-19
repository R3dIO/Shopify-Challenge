package models

import "gorm.io/gorm"

type Item struct {
	gorm.Model
	Name        string `json:"name"`
	City        string `json:"city"`
	Description string `json:"description"`
	Quantity	int	   `json:"quantity" gorm:"default:1"`
}

func GetOneItem(id string, item *Item) error {
	err := DB.Where("id = ?", id).First(&item).Error;
	return err
}

func CreateItem(item *Item) error {
	err := DB.Create(&item).Error
	return err
}

func UpdateItem(item *Item, input *Item) error {
	err := DB.Model(&item).Updates(&input).Error
	return err
}

func DeleteItem(item *Item) error {
	err := DB.Delete(&item).Error
	return err
}