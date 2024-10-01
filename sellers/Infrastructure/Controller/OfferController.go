package controller

import (
	"net/http"
	model "sellers/Domain/Model"
	repository "sellers/Infrastructure/Repository"

	"github.com/gin-gonic/gin"
)

func GetOfferByCardId(c *gin.Context) {
	cardId := c.Param("cardId")

	offer, err := repository.FindByCardId(cardId)
	if err != nil {
		c.IndentedJSON(http.StatusNotFound, err.Error())
		return
	}

	c.IndentedJSON(http.StatusAccepted, offer)
}

func PostOffer(c *gin.Context) {
	var offer model.Offer

	if err := c.BindJSON(&offer); err != nil {
		c.IndentedJSON(http.StatusNotFound, "Offer should be provided")
		return
	}

	repository.SaveOffer(&offer)

	c.IndentedJSON(http.StatusOK, offer)
}
