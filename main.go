package main

import (
	"encoding/json"
	"fmt"
	"math"
	"net/http"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/google/uuid"
	"github.com/gorilla/mux"
)

type Response struct {
	Message string `json:"message"`
}

type Receipt struct {
	Retailer     string `json:"retailer"`
	PurchaseDate string `json:"purchaseDate"`
	PurchaseTime string `json:"purchaseTime"`
	Items        []Item `json:"items"`
	Total        string `json:"total"`
}

type Item struct {
	ShortDescription string `json:"shortDescription"`
	Price            string `json:"price"`
}

type processReceiptResponse struct {
	ID string `json:"id"`
}

type getPointsResponse struct {
	Points int `json:"points"`
}

func calculatePoints(r Receipt) int {
	points := 0
	for _, char := range r.Retailer {
		if (char >= 'a' && char <= 'z') || (char >= 'A' && char <= 'Z') || (char >= '0' && char <= '9') {
			points += 1
		}
	}
	total, _ := strconv.ParseFloat(r.Total, 64)
	if math.Mod(total, 1) == 0 {
		points += 50
	}
	if math.Mod(total, 0.25) == 0 {
		points += 25
	}
	points += 5 * (len(r.Items) / 2)
	for _, item := range r.Items {
		trimmed := strings.TrimSpace(item.ShortDescription)
		if len(trimmed)%3 == 0 {
			price, _ := strconv.ParseFloat(item.Price, 64)
			points += int(math.Ceil(price * 0.2))
		}
	}
	purchaseDate, _ := time.Parse("2006-01-02", r.PurchaseDate)
	purchaseTime, _ := time.Parse("15:04", r.PurchaseTime)
	if purchaseDate.Day()%2 == 1 {
		points += 6
	}
	if purchaseTime.Hour() >= 14 && purchaseTime.Hour() < 16 {
		points += 10
	}

	return points
}

func processReceiptsHandler(w http.ResponseWriter, r *http.Request) {
	var receipt Receipt
	err := json.NewDecoder(r.Body).Decode(&receipt)
	if err != nil {
		http.Error(w, "Invalid receipt", http.StatusBadRequest)
		return
	}
	id := uuid.New().String()
	mutex.Lock()
	receipts[id] = receipt
	mutex.Unlock()
	response := processReceiptResponse{ID: id}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response)
}

func getPointsHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	mutex.Lock()
	receipt, exists := receipts[id]
	mutex.Unlock()
	if !exists {
		http.Error(w, "Receipt not found", http.StatusNotFound)
		return
	}
	points := calculatePoints(receipt)
	response := getPointsResponse{Points: points}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response)
}

var receipts = make(map[string]Receipt)
var mutex = &sync.Mutex{}

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/receipts/process", processReceiptsHandler).Methods("POST")
	r.HandleFunc("/receipts/{id}/points", getPointsHandler).Methods("GET")
	fmt.Println("Server is running on :8080")
	http.ListenAndServe(":8080", r)
}
