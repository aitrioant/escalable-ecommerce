package main

import (
	"fmt"
	"net/http"

	"database/sql"

	"github.com/gin-gonic/gin"
)

var db *sql.DB

func helloWorld(c *gin.Context) {
	fmt.Println("Hello World!")
	c.IndentedJSON(http.StatusOK, "hello sellers")
}

func main() {
	db = connectDB()
	defer db.Close()

	router := gin.Default()

	router.GET("/sellers", helloWorld)

	router.Run("localhost:8081")
}

func connectDB() *sql.DB {
	// Replace "username", "password", "dbname" with your database credentials
	connectionString := "username:password@tcp(localhost:3306)/sellers"
	db, err := sql.Open("mysql", connectionString)
	if err != nil {
		fmt.Println(err)
	}
	return db
}
