package main

import (
	"log"

	"receipt-processor-challenge/internal/handler"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	router.GET("/receipts/:id/points", handler.GetPoints)
	router.POST("/receipts/process", handler.ProcessReceipts)

	log.Println("Listening on localhost:8080")
	router.Run()
}
