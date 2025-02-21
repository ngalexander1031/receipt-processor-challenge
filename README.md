# receipt-processor-challenge

A Go-based API that processes receipt information and calculates reward points based on specific rules.

## Installation
Requires Go version 1.24.0

These packages need to be installed:
```bash
go install github.com/google/uuid
go install github.com/gorilla/mux
```

## How to run from root directory
### Start API
```bash
go run main.go
```
The server will start on port 8080.

### Run tests
```bash
go run test/test_api.go
```

## API Endpoints

### Process Receipt
**POST** `/receipts/process`
- Processes a receipt and returns a unique ID
- Request body should be a JSON object containing receipt details

### Get Points
**GET** `/receipts/{id}/points`
- Returns the points calculated for a processed receipt
- Uses the ID returned from the process endpoint

## Points Calculation Rules
The API calculates points based on the following rules:
1. One point for each alphanumeric character in the retailer name
2. 50 points if the total is a round dollar amount with no cents
3. 25 points if the total is a multiple of 0.25
4. 5 points for every two items on the receipt
5. If the trimmed length of an item description is a multiple of 3, multiply the price by 0.2 and round up to the nearest integer
6. 6 points if the purchase date is odd
7. 10 points if the purchase time is between 2:00pm and 4:00pm

## Data Storage
- Receipts are stored in-memory using a mutex-protected map
- Generated IDs use UUID

## Notes
- main.go was written without help from LLM.
- Test file was created with help from LLM and samples provided in original repo.
- README.md was formatted then improved with LLM.
- For additional API usage instructions, please refer to the [original challenge repository](https://github.com/fetch-rewards/receipt-processor-challenge?tab=readme-ov-file).