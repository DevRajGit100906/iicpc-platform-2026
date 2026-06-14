package main
import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"time"
	"github.com/gorilla/websocket"
	"github.com/redis/go-redis/v9"
)
var upgrader = websocket.Upgrader{CheckOrigin: func(r *http.Request) bool { return true }}
type Row struct {
	RunID       string  `json:"run_id"`
	Score       float64 `json:"score"`
	Status      string  `json:"status"`
	Correctness float64 `json:"correctness"`
	Passed      bool    `json:"passed"`
	Gated       bool    `json:"gated"`
	P99US       int64   `json:"p99_us"`
	TPS         float64 `json:"tps"`
}
func main() {
	rdb := redis.NewClient(&redis.Options{Addr: "localhost:6379"})
	ctx := context.Background()
	http.HandleFunc("/ws", func(w http.ResponseWriter, r *http.Request) {
		conn, err := upgrader.Upgrade(w, r, nil)
		if err != nil {
			return
		}
		defer conn.Close()
		fmt.Println("Frontend connected.")
		for {
			vals, err := rdb.ZRevRangeWithScores(ctx, "live_leaderboard", 0, 9).Result()
			if err != nil {
				time.Sleep(time.Second)
				continue
			}
			board := make([]Row, 0, len(vals))
			for _, z := range vals {
				id, _ := z.Member.(string)
				h, _ := rdb.HGetAll(ctx, "run:"+id).Result()
				row := Row{RunID: id, Score: z.Score, Status: h["status"]}
				if row.Status == "" {
					row.Status = "UNKNOWN"
				}
				row.Correctness = atof(h["correctness"])
				row.Passed = h["passed"] == "1"
				row.Gated = h["gated"] == "1"
				row.P99US = atoi(h["p99_us"])
				row.TPS = atof(h["tps"])
				board = append(board, row)
			}
			msg, _ := json.Marshal(board)
			if err := conn.WriteMessage(websocket.TextMessage, msg); err != nil {
				break
			}
			time.Sleep(500 * time.Millisecond)
		}
	})
	fmt.Println("Leaderboard WebSocket server on :8081")
	if err := http.ListenAndServe(":8081", nil); err != nil {
		panic(err)
	}
}
func atof(s string) float64 { f, _ := strconv.ParseFloat(s, 64); return f }
func atoi(s string) int64   { i, _ := strconv.ParseInt(s, 10, 64); return i }
