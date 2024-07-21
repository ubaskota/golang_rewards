package main

import (
	"github.com/gin-gonic/gin"
)

var recordStore map[string]TransactionDetail
var recordPoints map[string]int64

func init() {
	recordStore = make(map[string]TransactionDetail)
	recordPoints = make(map[string]int64)
}

func main() {
	router := gin.Default()
	router.GET("/receipts/:id/points", getPointsById)
	router.POST("/receipts/process", processTransaction)
	router.Run("localhost:8080")
}
