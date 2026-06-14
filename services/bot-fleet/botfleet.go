package botfleet

import (
	"bytes"
	"context"
	"encoding/json"
	"math/rand"
	"net/http"
	"sync"
	"sync/atomic"
	"time"
)

type Order struct {
	SequenceNumber int64  `json:"sequence_number"`
	OrderType      string `json:"order_type"`
	Side           string `json:"side"`
	Quantity       int32  `json:"quantity"`
	Price          int64  `json:"price"`
}

type Ack struct {
	SequenceNumber int64  `json:"sequence_number"`
	Status         string `json:"status"`
	AckTs          int64  `json:"ack_ts"`
}

type Telemetry struct {
	RunID          string `json:"run_id"`
	SequenceNumber int64  `json:"sequence_number"`
	SendTs         int64  `json:"send_ts"`
	AckTs          int64  `json:"ack_ts"`
	LatencyMicro   int64  `json:"latency_us"`
	Status         string `json:"status"`
}

type Engine interface {
	Send(Order) (Ack, error)
}

type Sink interface {
	Emit(Telemetry) error
}

type Config struct {
	RunID        string
	Bots         int
	OrdersPerBot int
	RatePerBot   float64
	Seed         int64
}

type Result struct {
	Sent   int64
	Errors int64
}

func Run(ctx context.Context, eng Engine, sink Sink, cfg Config) Result {
	if cfg.Bots <= 0 {
		cfg.Bots = 1
	}
	var seq, sent, errs int64
	var wg sync.WaitGroup
	for b := 0; b < cfg.Bots; b++ {
		wg.Add(1)
		go func(botID int) {
			defer wg.Done()
			rng := rand.New(rand.NewSource(cfg.Seed + int64(botID)))
			var interval time.Duration
			if cfg.RatePerBot > 0 {
				interval = time.Duration(float64(time.Second) / cfg.RatePerBot)
			}
			next := time.Now()
			for k := 0; k < cfg.OrdersPerBot; k++ {
				select {
				case <-ctx.Done():
					return
				default:
				}
				if interval > 0 {
					if d := time.Until(next); d > 0 {
						time.Sleep(d)
					}
					next = next.Add(interval)
				}
				o := genOrder(rng, atomic.AddInt64(&seq, 1))
				send := time.Now()
				ack, err := eng.Send(o)
				recv := time.Now()
				if err != nil {
					atomic.AddInt64(&errs, 1)
					_ = sink.Emit(Telemetry{RunID: cfg.RunID, SequenceNumber: o.SequenceNumber,
						SendTs: send.UnixNano(), Status: "error"})
					continue
				}
				atomic.AddInt64(&sent, 1)
				_ = sink.Emit(Telemetry{
					RunID:          cfg.RunID,
					SequenceNumber: o.SequenceNumber,
					SendTs:         send.UnixNano(),
					AckTs:          recv.UnixNano(),
					LatencyMicro:   recv.Sub(send).Microseconds(),
					Status:         ack.Status,
				})
			}
		}(b)
	}
	wg.Wait()
	return Result{Sent: sent, Errors: errs}
}

func genOrder(rng *rand.Rand, seq int64) Order {
	side := "BUY"
	if rng.Intn(2) == 0 {
		side = "SELL"
	}
	typ := "LIMIT"
	if rng.Float64() < 0.15 {
		typ = "MARKET"
	}
	const mid = int64(50000)
	return Order{
		SequenceNumber: seq,
		OrderType:      typ,
		Side:           side,
		Price:          mid + int64(rng.Intn(21)-10),
		Quantity:       int32((rng.Intn(10) + 1) * 10),
	}
}

type HTTPEngine struct {
	url    string
	client *http.Client
}

func NewHTTPEngine(orderURL string) *HTTPEngine {
	return &HTTPEngine{url: orderURL, client: &http.Client{Timeout: 2 * time.Second}}
}

func (e *HTTPEngine) Send(o Order) (Ack, error) {
	body, _ := json.Marshal(o)
	resp, err := e.client.Post(e.url, "application/json", bytes.NewReader(body))
	if err != nil {
		return Ack{}, err
	}
	defer resp.Body.Close()
	var a Ack
	_ = json.NewDecoder(resp.Body).Decode(&a)
	return a, nil
}
