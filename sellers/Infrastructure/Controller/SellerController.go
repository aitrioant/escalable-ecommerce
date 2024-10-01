package controller

import (
	"net/http"
	model "sellers/Domain/Model"
	repository "sellers/Infrastructure/Repository"

	"github.com/gin-gonic/gin"
)

func GetSellerById(c *gin.Context) {
	id := c.Param("id")

	seller, err := repository.FindSeller(id)
	if err != nil {
		c.IndentedJSON(http.StatusNotFound, err.Error())
		return
	}

	c.IndentedJSON(http.StatusAccepted, seller)
}

func PostSeller(c *gin.Context) {
	var seller model.Seller

	if err := c.BindJSON(&seller); err != nil {
		c.IndentedJSON(http.StatusNotFound, "Seller should be provided")
		return
	}

	repository.SaveSeller(&seller)

	c.IndentedJSON(http.StatusOK, seller)
}
