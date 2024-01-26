package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// User represents a user with a balance
type User struct {
	ID      string
	Name    string
	Balance float64
}

// Database storing user balances
var userDB = map[string]User{
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