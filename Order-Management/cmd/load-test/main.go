// load-test sends concurrent order requests to demonstrate race-safe stock reservations.
// Example: go run ./cmd/load-test -product-id 1 -customer-id 1 -quantity 1 -requests 100
package main

import (
	"bytes"
	"encoding/json"
	"flag"
	"fmt"
	"net/http"
	"sync"
	"sync/atomic"
	"time"
)

type requestBody struct {
	CustomerID uint `json:"customerId"`
	ProductID  uint `json:"productId"`
	Quantity   int  `json:"quantity"`
}

func main() {
	baseURL := flag.String("url", "http://localhost:8082", "Order Service base URL")
	productID := flag.Uint("product-id", 1, "Existing product ID")
	customerID := flag.Uint("customer-id", 1, "Customer ID")
	quantity := flag.Int("quantity", 1, "Quantity per order")
	requests := flag.Int("requests", 100, "Number of concurrent requests")
	flag.Parse()
	if *productID == 0 || *customerID == 0 || *quantity <= 0 || *requests <= 0 {
		fmt.Println("product-id, customer-id, quantity and requests must be positive")
		return
	}

	client := &http.Client{Timeout: 10 * time.Second}
	var wg sync.WaitGroup
	var succeeded, clientErrors, serverErrors, conflicts, other atomic.Int64
	start := time.Now()
	for i := 0; i < *requests; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			payload, _ := json.Marshal(requestBody{CustomerID: *customerID, ProductID: *productID, Quantity: *quantity})
			resp, err := client.Post(*baseURL+"/orders", "application/json", bytes.NewReader(payload))
			if err != nil {
				clientErrors.Add(1)
				return
			}
			defer resp.Body.Close()
			switch {
			case resp.StatusCode >= 200 && resp.StatusCode < 300:
				succeeded.Add(1)
			case resp.StatusCode == http.StatusConflict:
				conflicts.Add(1)
			case resp.StatusCode >= 500:
				serverErrors.Add(1)
			default:
				other.Add(1)
			}
		}()
	}
	wg.Wait()
	fmt.Printf("Concurrent order test finished in %s\n", time.Since(start))
	fmt.Printf("Requests: %d | Success: %d | Stock conflicts: %d | Server errors: %d | Client errors: %d | Other responses: %d\n", *requests, succeeded.Load(), conflicts.Load(), serverErrors.Load(), clientErrors.Load(), other.Load())
	fmt.Println("Stock conflicts are expected when requested total quantity exceeds available stock.")
}
