package main

import (
	"fmt"
	model "sellers/Domain/Model"
	controller "sellers/Infrastructure/Controller"
	"sellers/config"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func loadFixtures() {
	config.DB.AutoMigrate(&model.Seller{})
	config.DB.Create(model.SellersData)

}

func main() {
	config.DB = connectDB()

	// loadFixtures()
	router := gin.Default()

	router.GET("/sellers/:id", controller.GetSellerById)
	router.POST("/sellers", controller.PostSeller)
	router.GET("/sellers/offers/:cardId", controller.GetOfferByCardId)
	router.POST("/sellers/offers", controller.PostOffer)

	router.Run("localhost:8081")
}

func connectDB() *gorm.DB {
	dsn := "root:password@tcp(localhost:3306)/sellers"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Println(err)
	}
	return db
}
