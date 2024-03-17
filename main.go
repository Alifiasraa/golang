package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"math/rand"
	"net/http"
	"time"
)

const (
	jsonFilePath    = "status.json"
	updateInterval  = 15 * time.Second
	serverPort      = ":8080"
)

type Status struct {
	Water int `json:"water"`
	Wind  int `json:"wind"`
}

func generateRandomValue() int {
	return rand.Intn(100) + 1
}

func getStatusWater(value int) string {
	switch {
	case value < 5:
		return "aman"
	case value >= 6 && value <= 8:
		return "siaga"
	default:
		return "bahaya"
	}
}

func getStatusWind(value int) string {
	switch {
	case value < 6:
		return "aman"
	case value >= 7 && value <= 15:
		return "siaga"
	default:
		return "bahaya"
	}
}

func updateStatus() {
	for {
		waterValue := generateRandomValue()
		windValue := generateRandomValue()

		status := Status{
			Water: waterValue,
			Wind:  windValue,
		}

		statusJSON, err := json.MarshalIndent(status, "", "  ")
		if err != nil {
			fmt.Println("Error marshaling status:", err)
			return
		}

		err = ioutil.WriteFile(jsonFilePath, statusJSON, 0644)
		if err != nil {
			fmt.Println("Error writing to file:", err)
			return
		}

		waterStatus := getStatusWater(waterValue)
		windStatus := getStatusWind(windValue)

		fmt.Printf("Water: %d m (%s) | Wind: %d m/s (%s)\n", waterValue, waterStatus, windValue, windStatus)
		time.Sleep(updateInterval)
	}
}

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "Microservice status page")
	})

	fmt.Printf("Starting status updater on port %s...\n", serverPort)
	go updateStatus() 

	err := http.ListenAndServe(serverPort, nil)
	if err != nil {
		fmt.Println("Error starting server:", err)
	}
}
