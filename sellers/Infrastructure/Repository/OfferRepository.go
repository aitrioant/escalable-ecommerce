package repository

import (
	"errors"
	model "sellers/Domain/Model"
	"sellers/config"
)

func SaveOffer(offer *model.Offer) {
	config.DB.Save(offer)
}

func FindByCardId(cardId string) (*model.Offer, error) {
	var offer model.Offer
	result := config.DB.Where("card_Id = ?", cardId).First(&offer)

	if result.Error != nil {
		return nil, errors.New("Offer with cardId: " + cardId + " not found")
	}

	return &offer, nil
}
