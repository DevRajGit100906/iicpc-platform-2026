package main

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/gorilla/websocket"
	"github.com/redis/go-redis/v9"
)

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool { return true },
}

type LeaderboardRow struct {
	RunID  string  `json:"run_id"`
	Score  float64 `json:"score"`
	Status string  `json:"status"` // NEW: The Live Match State
}

func main() {
	rdb := redis.NewClient(&redis.Options{Addr: "localhost:6379"})
	ctx := context.Background()

	http.HandleFunc("/ws", func(w http.ResponseWriter, r *http.Request) {
		conn, err := upgrader.Upgrade(w, r, nil)
		if err != nil {
			fmt.Println("WS Upgrade Error:", err)
			return
		}
		defer conn.Close()

		fmt.Println("Frontend connected to WebSocket!")

		for {
			// Fetch the top 10 scores
			vals, err := rdb.ZRevRangeWithScores(ctx, "live_leaderboard", 0, 9).Result()
			if err != nil {
				time.Sleep(1 * time.Second)
				continue
			}

			var board []LeaderboardRow
			for _, z := range vals {
				runID := z.Member.(string)

				// NEW: Fetch the live state from the Match State Machine
				status, _ := rdb.HGet(ctx, "match:"+runID, "status").Result()
				if status == "" {
					status = "UNKNOWN"
				}

				board = append(board, LeaderboardRow{
					RunID:  runID,
					Score:  z.Score,
					Status: status,
				})
			}

			msg, _ := json.Marshal(board)
			if err := conn.WriteMessage(websocket.TextMessage, msg); err != nil {
				fmt.Println("Frontend disconnected.")
				break
			}
			time.Sleep(500 * time.Millisecond)
		}
	})

	fmt.Println("Leaderboard WebSocket Server listening on :8081...")
	if err := http.ListenAndServe(":8081", nil); err != nil {
		panic(err)
	}
}
