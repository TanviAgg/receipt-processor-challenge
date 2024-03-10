package handler

import (
	"fmt"
	"log"
	"net/http"
	"receipt-processor-challenge/internal/api"
	"receipt-processor-challenge/internal/models"
	"sync"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

var db sync.Map

func PrintSyncMap(m sync.Map) {
	// print map,
	fmt.Println("map content:")
	i := 0
	m.Range(func(key, value interface{}) bool {
		fmt.Printf("\t[%d] key: %v, value: %v\n", i, key, value)
		i++
		return true
	})
}

func ProcessReceipts(c *gin.Context) {
	var purchase models.Purchase

	err := c.ShouldBindJSON(&purchase)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	points, err := api.CalculatePoints(&purchase)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	id := uuid.NewString()

	db.Store(id, points)
	// PrintSyncMap(db)

	c.JSON(http.StatusCreated, gin.H{"id": id})
}

func GetPoints(c *gin.Context) {
	id := c.Param("id")
	log.Println("id", id)
	// PrintSyncMap(db)
	points, found := db.Load(id)
	log.Println("points", points)
	if !found {
		c.JSON(http.StatusBadRequest, gin.H{"error": "id not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"points": points})
}
