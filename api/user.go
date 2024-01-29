package api

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/jacobbeck/currency-converter/pkg/user"
)

// Database storing user balances
var userDB = map[string] user.User{
	"1": {ID: "1", Name: "Alice", Balance: 100.0},
	"2": {ID: "2", Name: "Bob", Balance: 50.0},
	// Add more users as needed
}

// getUserBalance returns the balance of a user given their ID
func GetUserBalance(userID string) (float64, bool) {
	user, ok := userDB[userID]
	if !ok {
		return 0, false
	}
	return user.Balance, true
}

// getUserBalanceHandler is a handler function to get a user's balance
func GetUserBalanceHandler(c *gin.Context) {
	userID := c.Param("id")
	balance, ok := GetUserBalance(userID)
	if !ok {
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"balance": balance})
}

func CreateUserHandler(c *gin.Context) {
	userService, exists := c.Get("userService")
	if !exists {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Invalid user type in context"})
		return
	}

	us, _ := userService.(user.Service)

	var u user.User
	if err := c.BindJSON(&u); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	
	newUser, err := us.CreateUser(&u)
	if err != nil {
		fmt.Println(err)
		c.JSON(http.StatusInternalServerError, err)
	}

	// Return the newly created user
	c.JSON(http.StatusCreated, newUser)
}


func UserMiddleware(userService user.Service) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Set("userService", userService)
		c.Next()
	}
}