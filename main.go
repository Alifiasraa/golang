package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"golang-fga/database"
	"golang-fga/models"

	"gorm.io/gorm"
)

func main() {
	database.StartDB()
	// createOrder("kazuha")
	// getOrderById(1)
	getAllOrders()
	// updateOrderById(1, "childe")
	// deleteOrderById(2)
}

func createOrder(name string) {
	db := database.GetDB()

	order := models.Order{
		CustomerName: name,
	}

	err := db.Debug().Create(&order).Error
	if err != nil {
		fmt.Println("Failed to create new order, err:", err)
		return
	}

	fmt.Println("Success to create new order:", order)
}

func getAllOrders() {
    db := database.GetDB()

    var orders []models.Order

    err := db.Debug().Order("id").Preload("Items").Find(&orders).Error
    if err != nil {
        fmt.Println("Failed to fetch orders, err:", err)
        return
    }

	// convert to JSON
    ordersJSON, err := json.Marshal(orders)
    if err != nil {
        fmt.Println("Failed to marshal orders to JSON, err:", err)
        return
    }

    fmt.Println("Orders Data:", string(ordersJSON))
}

func getOrderById(id uint) {
	db := database.GetDB()

	order := models.Order{}

	err := db.Debug().First(&order, "id=?", id).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			fmt.Println("Order data not found")
			return
		}

		fmt.Println("Failed to find order data, err:", err)
		return
	}

	fmt.Printf("Order data: %+v\n", order)
}

func updateOrderById(id uint, name string) {
	db := database.GetDB()

	order := models.Order{}

	err := db.Model(&order).Where("id=?", id).Updates(models.Order{CustomerName: name}).Error
	if err != nil {
		fmt.Println("Failed to update order data, err:", err)
		return
	}

	fmt.Printf("Success to update order data: %+v\n", order)
}

func deleteOrderById(id uint) {
	db := database.GetDB()

	order := models.Order{}

	err := db.Where("id=?", id).Delete(&order).Error
	if err != nil {
		fmt.Println("Failed to delete order, err:", err)
		return
	}

	fmt.Printf("Order with id %d has been deleted", id)
}
