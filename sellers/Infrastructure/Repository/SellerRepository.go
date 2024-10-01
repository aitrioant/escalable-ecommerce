package repository

import (
	"errors"
	model "sellers/Domain/Model"
	"sellers/config"
)

func SaveSeller(seller *model.Seller) {
	config.DB.Save(seller)
}

func FindSeller(id string) (*model.Seller, error) {
	seller := model.Seller{ID: id}
	result := config.DB.First(&seller, id)

	if result.Error != nil {
		return nil, errors.New("Seller with ID: " + id + " not found")
	}

	return &seller, nil
}
