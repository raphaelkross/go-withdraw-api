package controllers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// Withdraw Request Controller
func Withdraw(c *gin.Context) {
	// Try converting amount to inmt
	amount, err := strconv.Atoi(c.Param("amount"))

	// Return error if non-integer value is provided.
	if err != nil {
		c.JSON(http.StatusBadRequest, map[string]string{"error": "Only integer amounts are accepted."})
		return
	}

	c.JSON(http.StatusOK, amount)
}
