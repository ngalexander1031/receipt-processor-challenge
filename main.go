package main

import (
	"fmt"
	"net/http"
	"sync"

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

type ProcessResponse struct {
	ID string `json:"id"`
}

type PointsResponse struct {
	Points int `json:"points"`
}

func calculatePoints() int {
	points := 0

	return points
}

func processReceiptsHandler(w http.ResponseWriter, r *http.Request) {

}

func getPointsHandler(w http.ResponseWriter, r *http.Request) {

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
