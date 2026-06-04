package main

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"

	"github.com/segmentio/kafka-go"
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
}

func main() {
	sutURL := "http://localhost:8080/order"
	client := &http.Client{Timeout: 2 * time.Second}
	runID := fmt.Sprintf("run-%d", time.Now().Unix())

	// Configure the Kafka Writer
	// Configure the Kafka Writer
	writer := &kafka.Writer{
		Addr:                   kafka.TCP("localhost:19092"),
		Topic:                  "telemetry",
		Balancer:               &kafka.LeastBytes{},
		AllowAutoTopicCreation: true, // <-- Add this
		MaxAttempts:            5,    // <-- Add this (gives Kafka time to elect a leader)
	}
	defer writer.Close()

	fmt.Printf("Starting Kafka-connected bot worker for %s...\n", runID)

	for i := int64(1); i <= 10; i++ {
		order := Order{
			SequenceNumber: i,
			OrderType:      "LIMIT",
			Side:           "BUY",
			Quantity:       100,
			Price:          50000,
		}

		payload, _ := json.Marshal(order)
		req, _ := http.NewRequest("POST", sutURL, bytes.NewBuffer(payload))
		req.Header.Set("Content-Type", "application/json")

		sendTs := time.Now().UnixNano()

		resp, err := client.Do(req)
		if err != nil {
			fmt.Printf("Failed to send order %d: %v\n", i, err)
			continue
		}

		var ack Ack
		body, _ := io.ReadAll(resp.Body)
		json.Unmarshal(body, &ack)
		resp.Body.Close()

		latencyMicro := (ack.AckTs - sendTs) / time.Microsecond.Nanoseconds()

		t := Telemetry{
			RunID:          runID,
			SequenceNumber: i,
			SendTs:         sendTs,
			AckTs:          ack.AckTs,
			LatencyMicro:   latencyMicro,
		}

		// Serialize and push to Kafka
		telemetryBytes, _ := json.Marshal(t)
		err = writer.WriteMessages(context.Background(),
			kafka.Message{
				Key:   []byte(runID),
				Value: telemetryBytes,
			},
		)

		if err != nil {
			fmt.Printf("[ERROR] Failed to push telemetry to Kafka: %v\n", err)
		} else {
			fmt.Printf("[KAFKA] Pushed Seq: %d | Latency: %d µs\n", t.SequenceNumber, t.LatencyMicro)
		}

		time.Sleep(200 * time.Millisecond)
	}

	fmt.Println("Load generation and telemetry streaming complete.")
}
