package correctness
import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"sync"
	"testing"
	"time"
	"github.com/devrajdas/iicpc-platform-2026/pkg/orderbook"
)
func goodEngine() http.Handler {
	ob := orderbook.NewOrderBook()
	var mu sync.Mutex
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var o WireOrder
		json.NewDecoder(r.Body).Decode(&o)
		mu.Lock()
		resp := OrderResponse{SequenceNumber: o.SequenceNumber, Status: "ACCEPTED", AckTs: time.Now().UnixNano()}
		if o.OrderType == "CANCEL" {
			if ob.Cancel(o.OrigSequenceNumber) {
				resp.Status = "CANCELED"
			} else {
				resp.Status = "REJECTED"
			}
		} else {
			for _, t := range ob.Submit(toOB(o)) {
				resp.Fills = append(resp.Fills, Fill{Price: t.Price, Quantity: t.Quantity, MakerSeq: t.MakerID})
			}
			if len(resp.Fills) > 0 {
				resp.Status = "FILLED"
			}
		}
		mu.Unlock()
		json.NewEncoder(w).Encode(resp)
	})
}
func lazyEngine() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var o WireOrder
		json.NewDecoder(r.Body).Decode(&o)
		json.NewEncoder(w).Encode(OrderResponse{SequenceNumber: o.SequenceNumber, Status: "ACCEPTED", AckTs: time.Now().UnixNano()})
	})
}
func wrongPriceEngine() http.Handler {
	ob := orderbook.NewOrderBook()
	var mu sync.Mutex
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var o WireOrder
		json.NewDecoder(r.Body).Decode(&o)
		mu.Lock()
		resp := OrderResponse{SequenceNumber: o.SequenceNumber, Status: "ACCEPTED", AckTs: time.Now().UnixNano()}
		if o.OrderType != "CANCEL" {
			for _, t := range ob.Submit(toOB(o)) {
				resp.Fills = append(resp.Fills, Fill{Price: t.Price + 1, Quantity: t.Quantity, MakerSeq: t.MakerID})
			}
			if len(resp.Fills) > 0 {
				resp.Status = "FILLED"
			}
		}
		mu.Unlock()
		json.NewEncoder(w).Encode(resp)
	})
}
func TestGoodEnginePasses(t *testing.T) {
	srv := httptest.NewServer(goodEngine())
	defer srv.Close()
	orders := GenerateScenario(123, 500)
	rep, err := Check(orders, NewHTTPClient(srv.URL))
	if err != nil {
		t.Fatal(err)
	}
	if rep.ExpectedTrades == 0 {
		t.Fatal("scenario produced no trades — generator isn't crossing, test is meaningless")
	}
	if !rep.Passed {
		t.Fatalf("correct engine should pass; first violation: %+v", rep.FirstViolation)
	}
	if rep.Score != 1.0 {
		t.Fatalf("want score 1.0, got %.4f", rep.Score)
	}
	t.Logf("OK: %d orders -> %d trades, score %.2f", rep.TotalOrders, rep.ExpectedTrades, rep.Score)
}
func TestLazyEngineCaught(t *testing.T) {
	srv := httptest.NewServer(lazyEngine())
	defer srv.Close()
	rep, _ := Check(GenerateScenario(123, 500), NewHTTPClient(srv.URL))
	if rep.Passed {
		t.Fatal("an engine that never matches must NOT pass")
	}
	t.Logf("caught no-fill engine: expected %d trades, got %d (score %.2f)", rep.ExpectedTrades, rep.GotTrades, rep.Score)
}
func TestWrongPriceCaughtAndLocalized(t *testing.T) {
	srv := httptest.NewServer(wrongPriceEngine())
	defer srv.Close()
	rep, _ := Check(GenerateScenario(123, 500), NewHTTPClient(srv.URL))
	if rep.Passed {
		t.Fatal("a wrong-price engine must NOT pass")
	}
	if rep.FirstViolation == nil || rep.FirstViolation.Note != "" {
		t.Fatalf("should localize a value mismatch, got %+v", rep.FirstViolation)
	}
	v := rep.FirstViolation
	if v.Got.Price != v.Expected.Price+1 {
		t.Fatalf("expected the +1 tick bug to be the divergence, got expected=%+v got=%+v", v.Expected, v.Got)
	}
	t.Logf("localized at trade #%d (order seq %d): expected price %d, engine reported %d",
		v.Index, v.TakerSeq, v.Expected.Price, v.Got.Price)
}
