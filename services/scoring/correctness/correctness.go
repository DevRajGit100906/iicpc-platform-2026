package correctness
import (
	"bytes"
	"encoding/json"
	"math/rand"
	"net/http"
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
	Fills          []Fill `json:"fills"`
}
type EngineClient interface {
	Submit(WireOrder) (OrderResponse, error)
}
type HTTPClient struct {
	url string
	c   *http.Client
}
func NewHTTPClient(orderURL string) *HTTPClient {
	return &HTTPClient{url: orderURL, c: &http.Client{Timeout: 5 * time.Second}}
}
func (h *HTTPClient) Submit(o WireOrder) (OrderResponse, error) {
	b, _ := json.Marshal(o)
	resp, err := h.c.Post(h.url, "application/json", bytes.NewReader(b))
	if err != nil {
		return OrderResponse{}, err
	}
	defer resp.Body.Close()
	var or OrderResponse
	if err := json.NewDecoder(resp.Body).Decode(&or); err != nil {
		return OrderResponse{}, err
	}
	return or, nil
}
type Trade struct {
	Taker    int64
	Maker    int64
	Price    int64
	Quantity int64
}
type Violation struct {
	Index    int
	TakerSeq int64
	Expected Trade
	Got      Trade
	Note     string
}
type Report struct {
	TotalOrders    int
	ExpectedTrades int
	GotTrades      int
	Matched        int
	Score          float64
	Passed         bool
	FirstViolation *Violation
}
func GenerateScenario(seed int64, n int) []WireOrder {
	rng := rand.New(rand.NewSource(seed))
	const mid = int64(50000)
	orders := make([]WireOrder, 0, n)
	for i := int64(1); i <= int64(n); i++ {
		o := WireOrder{SequenceNumber: i, Side: "BUY", OrderType: "LIMIT", Quantity: int32((rng.Intn(5) + 1) * 10)}
		if rng.Intn(2) == 0 {
			o.Side = "SELL"
		}
		if rng.Float64() < 0.2 {
			o.OrderType = "MARKET"
		} else {
			o.Price = mid + int64(rng.Intn(11)-5)
		}
		orders = append(orders, o)
	}
	return orders
}
func toOB(o WireOrder) *orderbook.Order {
	return &orderbook.Order{
		ID:       o.SequenceNumber,
		Side:     orderbook.Side(o.Side),
		Type:     orderbook.OrderType(o.OrderType),
		Price:    o.Price,
		Quantity: int64(o.Quantity),
		TIF:      orderbook.GTC,
	}
}
func oracleTrades(orders []WireOrder) []Trade {
	ob := orderbook.NewOrderBook()
	var out []Trade
	for _, o := range orders {
		if o.OrderType == "CANCEL" {
			ob.Cancel(o.OrigSequenceNumber)
			continue
		}
		for _, t := range ob.Submit(toOB(o)) {
			out = append(out, Trade{Taker: t.TakerID, Maker: t.MakerID, Price: t.Price, Quantity: t.Quantity})
		}
	}
	return out
}
func engineTrades(orders []WireOrder, client EngineClient) ([]Trade, error) {
	var out []Trade
	for _, o := range orders {
		resp, err := client.Submit(o)
		if err != nil {
			return nil, err
		}
		for _, f := range resp.Fills {
			out = append(out, Trade{Taker: o.SequenceNumber, Maker: f.MakerSeq, Price: f.Price, Quantity: f.Quantity})
		}
	}
	return out, nil
}
func Check(orders []WireOrder, client EngineClient) (Report, error) {
	got, err := engineTrades(orders, client)
	if err != nil {
		return Report{}, err
	}
	want := oracleTrades(orders)
	r := Report{TotalOrders: len(orders), ExpectedTrades: len(want), GotTrades: len(got)}
	n := len(want)
	if len(got) < n {
		n = len(got)
	}
	for i := 0; i < n; i++ {
		if want[i] == got[i] {
			r.Matched++
			continue
		}
		r.FirstViolation = &Violation{Index: i, TakerSeq: want[i].Taker, Expected: want[i], Got: got[i]}
		break
	}
	if r.FirstViolation == nil && len(want) != len(got) {
		r.FirstViolation = &Violation{Index: n, Note: "trade-count mismatch: engine produced a different number of fills than the oracle"}
	}
	if len(want) == 0 {
		r.Score = 1.0
	} else {
		r.Score = float64(r.Matched) / float64(len(want))
	}
	r.Passed = r.FirstViolation == nil && len(want) == len(got)
	return r, nil
}