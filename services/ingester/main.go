package main

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	"github.com/redis/go-redis/v9"
	"github.com/segmentio/kafka-go"
)

type Telemetry struct {
	RunID          string `json:"run_id"`
	SequenceNumber int64  `json:"sequence_number"`
	SendTs         int64  `json:"send_ts"`
	AckTs          int64  `json:"ack_ts"`
	LatencyMicro   int64  `json:"latency_us"`
}

func main() {
	fmt.Println("Starting Minimal Telemetry Ingester...")

	// 1. Connect to Redis (Used for the live leaderboard ZSET)
	rdb := redis.NewClient(&redis.Options{
		Addr: "localhost:6379",
	})
	ctx := context.Background()
	if err := rdb.Ping(ctx).Err(); err != nil {
		fmt.Printf("Redis not ready: %v\n", err)
	} else {
		fmt.Println("Connected to Redis.")
	}

	// 2. Connect to Kafka
	reader := kafka.NewReader(kafka.ReaderConfig{
		Brokers:   []string{"localhost:19092"},
		Topic:     "telemetry",
		Partition: 0,
		MaxBytes:  10e6, // 10MB
	})
	defer reader.Close()

	fmt.Println("Listening for telemetry on Kafka topic 'telemetry'...")

	var totalOrders int64
	var maxLatency int64

	// 3. Consume Loop
	for {
		m, err := reader.ReadMessage(ctx)
		if err != nil {
			fmt.Printf("Kafka read error (waiting for broker): %v\n", err)
			time.Sleep(2 * time.Second)
			continue
		}

		var t Telemetry
		json.Unmarshal(m.Value, &t)

		totalOrders++
		if t.LatencyMicro > maxLatency {
			maxLatency = t.LatencyMicro
		}

		// 4. Update the Redis ZSET (The Leaderboard)
		// We use the RunID as the member, and an arbitrary composite score for now.
		// In Epic E, this will become the rigorous percentile-based score.
		score := float64(totalOrders) + (10000.0 / float64(maxLatency))

		err = rdb.ZAdd(ctx, "live_leaderboard", redis.Z{
			Score:  score,
			Member: t.RunID,
		}).Err()

		if err != nil {
			fmt.Printf("Failed to update Redis: %v\n", err)
		}

		fmt.Printf("[INGESTED] Run: %s | Total: %d | Max Latency: %d µs | Score: %.2f\n",
			t.RunID, totalOrders, maxLatency, score)
	}
}
