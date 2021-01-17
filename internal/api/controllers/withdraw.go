package controllers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/raphaelkross/go-withdraw-api/internal/helpers"
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

	output := helpers.SplitAmountIntoBills(amount)

	c.JSON(http.StatusOK, output)
}
