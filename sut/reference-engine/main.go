package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"sync"
	"time"

	"github.com/devrajdas/iicpc-platform-2026/pkg/orderbook"
)

type WireOrder struct {
	SequenceNumber     int64  `json:"sequence_number"`
	OrderType          string `json:"order_type"`
	Side               string `json:"side"`
	Quantity           int32  `json:"quantity"`
	Price              int64  `json:"price"`
	OrigSequenceNumber int64  `json:"orig_sequence_number,omitempty"`
}

type Fill struct {
	Price    int64 `json:"price"`
	Quantity int64 `json:"quantity"`
	MakerSeq int64 `json:"maker_seq"`
}

type OrderResponse struct {
	SequenceNumber int64  `json:"sequence_number"`
	Status         string `json:"status"`
	AckTs          int64  `json:"ack_ts"`
	Fills          []Fill `json:"fills,omitempty"`
}

type Engine struct {
	mu   sync.Mutex
	book *orderbook.OrderBook
}

func (e *Engine) handleOrder(w http.ResponseWriter, r *http.Request) {
	var o WireOrder
	if err := json.NewDecoder(r.Body).Decode(&o); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	e.mu.Lock()
	resp := OrderResponse{SequenceNumber: o.SequenceNumber, Status: "ACCEPTED", AckTs: time.Now().UnixNano()}
	if o.OrderType == "CANCEL" {
		if e.book.Cancel(o.OrigSequenceNumber) {
			resp.Status = "CANCELED"
		} else {
			resp.Status = "REJECTED"
		}
	} else {
		trades := e.book.Submit(&orderbook.Order{
			ID:       o.SequenceNumber,
			Side:     orderbook.Side(o.Side),
			Type:     orderbook.OrderType(o.OrderType),
			Price:    o.Price,
			Quantity: int64(o.Quantity),
			TIF:      orderbook.GTC,
		})
		for _, t := range trades {
			resp.Fills = append(resp.Fills, Fill{Price: t.Price, Quantity: t.Quantity, MakerSeq: t.MakerID})
		}
		if len(resp.Fills) > 0 {
			resp.Status = "FILLED"
		}
	}
	e.mu.Unlock()
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(resp)
}

func (e *Engine) handleHealth(w http.ResponseWriter, _ *http.Request) {
	w.WriteHeader(http.StatusOK)
	fmt.Fprint(w, "ok")
}

func main() {
	e := &Engine{book: orderbook.NewOrderBook()}
	http.HandleFunc("/order", e.handleOrder)
	http.HandleFunc("/healthz", e.handleHealth)
	fmt.Println("Reference matching engine on :8080 (price-time priority, reports fills)")
	fmt.Println("Yaay completed my matching engine , now lets test its correctness :)")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		panic(err)
	}
}
