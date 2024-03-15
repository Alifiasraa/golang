package main

import (
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// Define models
type Order struct {
	ID           uint      `gorm:"primaryKey"`
	CustomerName string    `json:"customerName"`
	OrderedAt    time.Time `json:"orderedAt"`
	Items        []Item    `json:"items" gorm:"foreignKey:OrderID"`
}

type Item struct {
	ID          uint   `gorm:"primaryKey"`
	Code        string `json:"itemCode"`
	Description string `json:"description"`
	Quantity    uint   `json:"quantity"`
	OrderID     uint   `json:"-"`
}

var (
	db *gorm.DB
)

func main() {
	// Connect to the database
	dsn := "host=localhost user=postgres password=yeay123 dbname=postgres port=5432 sslmode=disable"
	var err error
	db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("Failed to connect to database!")
	}

	r := gin.Default()

	// Define routes
	r.POST("/orders", createOrder)
	r.GET("/orders", getOrders)
	r.PUT("/orders/:id", updateOrder)
	r.DELETE("/orders/:id", deleteOrder)

	r.Run(":8080")
}

// Handlers
func createOrder(c *gin.Context) {
	var orderData Order

	if err := c.ShouldBindJSON(&orderData); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	db.Create(&orderData)

	c.JSON(201, orderData)
}

func getOrders(c *gin.Context) {
	var orders []Order
	db.Preload("Items").Find(&orders)

	c.JSON(200, orders)
}

func updateOrder(c *gin.Context) {
	var orderData Order
	id := c.Param("id")

	if err := db.Preload("Items").First(&orderData, id).Error; err != nil {
		c.JSON(404, gin.H{"error": "Record not found!"})
		return
	}

	var newData struct {
		OrderedAt    time.Time `json:"orderedAt"`
		CustomerName string    `json:"customerName"`
		Items        []struct {
			ID          uint   `json:"id"`
			ItemCode    string `json:"itemCode"`
			Description string `json:"description"`
			Quantity    uint   `json:"quantity"`
		} `json:"items"`
	}

	if err := c.ShouldBindJSON(&newData); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	orderData.OrderedAt = newData.OrderedAt
	orderData.CustomerName = newData.CustomerName

	for _, newItem := range newData.Items {
		existingItem := Item{
			ID:      newItem.ID,
			OrderID: orderData.ID,
		}
		if err := db.First(&existingItem).Error; err != nil {
			continue
		}

		existingItem.Code = newItem.ItemCode
		existingItem.Description = newItem.Description
		existingItem.Quantity = newItem.Quantity
		db.Save(&existingItem)
	}

	db.Save(&orderData)

	c.JSON(200, newData)
}

func deleteOrder(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(400, gin.H{"error": "Invalid ID format"})
		return
	}

	order := &Order{
		ID: uint(id),
	}

	if err := db.First(order).Error; err != nil {
		c.JSON(404, gin.H{"error": "Record not found!"})
		return
	}

	if err := db.Delete(order).Error; err != nil {
		c.JSON(500, gin.H{"error": "Failed delete data!"})
		return
	}

	c.JSON(200, gin.H{"message": "Success delete"})
}
