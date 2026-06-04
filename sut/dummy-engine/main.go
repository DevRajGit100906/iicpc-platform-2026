package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

// Order mimics the JSON payload expected from the Load Generator's REST adapter
type Order struct {
	SequenceNumber int64  `json:"sequence_number"`
	OrderType      string `json:"order_type"`
	Side           string `json:"side"`
	Quantity       int32  `json:"quantity"`
	Price          int64  `json:"price"`
}

// Ack is the response sent back to the Load Generator
type Ack struct {
	SequenceNumber int64  `json:"sequence_number"`
	Status         string `json:"status"`
	AckTs          int64  `json:"ack_ts"`
}

func handleOrder(w http.ResponseWriter, r *http.Request) {
	var req Order
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Simulate the latency of a baseline matching engine (e.g., 50 microseconds)
	time.Sleep(50 * time.Microsecond)

	ack := Ack{
		SequenceNumber: req.SequenceNumber,
		Status:         "ACCEPTED",
		AckTs:          time.Now().UnixNano(),
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(ack)
}

func main() {
	http.HandleFunc("/order", handleOrder)
	fmt.Println("Dummy SUT Matching Engine listening on :8080...")
	
	// In a real SUT, this would be highly optimized TCP/WebSockets, 
	// but REST is perfectly fine for establishing the Vertical Slice spine.
	if err := http.ListenAndServe(":8080", nil); err != nil {
		panic(err)
	}
}