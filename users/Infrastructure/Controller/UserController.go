package controller

import (
	"net/http"
	user "users/Domain/Model"
	"users/config"

	"github.com/gin-gonic/gin"
)

// getUsers responds with the list of all users as JSON.
func GetUsers(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, config.UserRepo.GetUsers())
}

// postUsers adds an user from JSON received in the request body.
func PostUsers(c *gin.Context) {
	var newUser user.User

	// Call BindJSON to bind the received JSON to
	// newuser.
	if err := c.BindJSON(&newUser); err != nil {
		return
	}

	// Add the new user to the slice.
	config.UserRepo.UpdateUser(newUser)

	c.IndentedJSON(http.StatusCreated, newUser)
}

// getUserByID locates the user whose ID value matches the id
// parameter sent by the client, then returns that user as a response.
func GetUserByID(c *gin.Context) {
	id := c.Param("id")

	// Loop through the list of users, looking for
	// an user whose ID value matches the parameter.
	a, error := config.UserRepo.GetUser(id)
	if !error {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "user not found"})
	}

	c.IndentedJSON(http.StatusOK, a)

}
