package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func processTransaction(c *gin.Context) {
	var transaction TransactionDetail
	if err := c.BindJSON(&transaction); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request data"})
		return
	}
	var transaction_id string = generateTransactionID()
	recordStore[transaction_id] = transaction
	var totalPoints int64 = calculateTotalPoints(transaction)
	fmt.Println("This is the transaction_id : ", transaction_id)
	recordPoints[transaction_id] = totalPoints
	data := IdResponse{ID: transaction_id}
	c.JSON(http.StatusOK, data)
}

func getPointsById(c *gin.Context) {
	id := c.Param("id")
	value, exists := recordPoints[id]
	if exists {
		data := PointsResponse{Points: value}
		c.JSON(http.StatusOK, data)
	} else {
		c.JSON(http.StatusNotFound, gin.H{"error": "Points not found for the given ID"})
	}
}

func generateTransactionID() string {
	return uuid.New().String()
}
