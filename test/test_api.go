package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type ProcessResponse struct {
	ID string `json:"id"`
}

type PointsResponse struct {
	Points int `json:"points"`
}

func main() {
	baseURL := "http://localhost:8080"

	receipts := []map[string]interface{}{
		{
			"retailer":     "Walgreens",
			"purchaseDate": "2022-01-02",
			"purchaseTime": "08:13",
			"total":        "2.65",
			"items": []map[string]string{
				{"shortDescription": "Pepsi - 12-oz", "price": "1.25"},
				{"shortDescription": "Dasani", "price": "1.40"},
			},
		},
		{
			"retailer":     "Target",
			"purchaseDate": "2022-01-02",
			"purchaseTime": "13:13",
			"total":        "1.25",
			"items": []map[string]string{
				{"shortDescription": "Pepsi - 12-oz", "price": "1.25"},
			},
		},
	}

	var ids []string

	// Process each receipt
	for _, receipt := range receipts {
		body, _ := json.Marshal(receipt)
		resp, err := http.Post(baseURL+"/receipts/process", "application/json", bytes.NewBuffer(body))
		if err != nil {
			fmt.Println("Error processing receipt:", err)
			return
		}
		defer resp.Body.Close()

		var result ProcessResponse
		respBody, _ := io.ReadAll(resp.Body)
		json.Unmarshal(respBody, &result)
		ids = append(ids, result.ID)

		fmt.Printf("Processed receipt: %s\n", result.ID)
	}

	// Retrieve points for each receipt
	for _, id := range ids {
		resp, err := http.Get(baseURL + "/receipts/" + id + "/points")
		if err != nil {
			fmt.Println("Error getting points:", err)
			return
		}
		defer resp.Body.Close()

		var result PointsResponse
		respBody, _ := io.ReadAll(resp.Body)
		json.Unmarshal(respBody, &result)

		fmt.Printf("Receipt ID: %s - Points: %d\n", id, result.Points)
	}
}
