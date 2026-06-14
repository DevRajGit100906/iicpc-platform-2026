package orderbook
import "testing"
func lim(ob *OrderBook, id int64, side Side, px, qty int64) []Trade {
	return ob.Submit(&Order{ID: id, Side: side, Type: Limit, Price: px, Quantity: qty, TIF: GTC})
}
func TestTimePriorityFIFO(t *testing.T) {
	ob := NewOrderBook()
	lim(ob, 1, Buy, 50000, 100)
	lim(ob, 2, Buy, 50000, 100)
	trades := lim(ob, 3, Sell, 50000, 150)
	if len(trades) != 2 {
		t.Fatalf("want 2 trades, got %d", len(trades))
	}
	if trades[0].MakerID != 1 || trades[0].Quantity != 100 {
		t.Errorf("first fill should be maker 1 for 100, got maker %d qty %d", trades[0].MakerID, trades[0].Quantity)
	}
	if trades[1].MakerID != 2 || trades[1].Quantity != 50 {
		t.Errorf("second fill should be maker 2 for 50, got maker %d qty %d", trades[1].MakerID, trades[1].Quantity)
	}
	if q := ob.RestingQty(Buy, 50000); q != 50 {
		t.Errorf("maker 2 should have 50 left resting, got %d", q)
	}
}
func TestPricePriority(t *testing.T) {
	ob := NewOrderBook()
	lim(ob, 1, Sell, 50001, 100)
	lim(ob, 2, Sell, 50000, 100)
	trades := lim(ob, 3, Buy, 50001, 100)
	if len(trades) != 1 {
		t.Fatalf("want 1 trade, got %d", len(trades))
	}
	if trades[0].MakerID != 2 || trades[0].Price != 50000 {
		t.Errorf("should hit better ask (maker 2 @ 50000), got maker %d @ %d", trades[0].MakerID, trades[0].Price)
	}
}
func TestNoTradeThrough(t *testing.T) {
	ob := NewOrderBook()
	lim(ob, 1, Sell, 50001, 100)
	trades := lim(ob, 2, Buy, 50000, 100)
	if len(trades) != 0 {
		t.Fatalf("want 0 trades, got %d", len(trades))
	}
	if bid, _ := ob.BestBid(); bid != 50000 {
		t.Errorf("buy should rest at 50000, best bid = %d", bid)
	}
}
func TestPartialFillRests(t *testing.T) {
	ob := NewOrderBook()
	lim(ob, 1, Buy, 50000, 100)
	lim(ob, 2, Sell, 50000, 60)
	if q := ob.RestingQty(Buy, 50000); q != 40 {
		t.Fatalf("want 40 resting, got %d", q)
	}
	trades := lim(ob, 3, Sell, 50000, 40)
	if len(trades) != 1 || trades[0].Quantity != 40 {
		t.Fatalf("want one 40-qty fill, got %+v", trades)
	}
	if q := ob.RestingQty(Buy, 50000); q != 0 {
		t.Errorf("book should be empty, got %d resting", q)
	}
}
func TestFOKReject(t *testing.T) {
	ob := NewOrderBook()
	lim(ob, 1, Sell, 50000, 50)
	trades := ob.Submit(&Order{ID: 2, Side: Buy, Type: Limit, Price: 50000, Quantity: 100, TIF: FOK})
	if len(trades) != 0 {
		t.Fatalf("FOK should reject with 0 trades, got %d", len(trades))
	}
	if q := ob.RestingQty(Sell, 50000); q != 50 {
		t.Errorf("resting sell should be untouched at 50, got %d", q)
	}
}
func TestFOKFills(t *testing.T) {
	ob := NewOrderBook()
	lim(ob, 1, Sell, 50000, 60)
	lim(ob, 2, Sell, 50001, 60)
	trades := ob.Submit(&Order{ID: 3, Side: Buy, Type: Limit, Price: 50001, Quantity: 100, TIF: FOK})
	var filled int64
	for _, tr := range trades {
		filled += tr.Quantity
	}
	if filled != 100 {
		t.Fatalf("FOK should fill exactly 100, filled %d", filled)
	}
}
func TestCancel(t *testing.T) {
	ob := NewOrderBook()
	lim(ob, 1, Buy, 50000, 100)
	if !ob.Cancel(1) {
		t.Fatal("cancel of resting order should succeed")
	}
	trades := lim(ob, 2, Sell, 50000, 100)
	if len(trades) != 0 {
		t.Errorf("cancelled order must not fill, got %d trades", len(trades))
	}
}
func TestIOCNoRest(t *testing.T) {
	ob := NewOrderBook()
	lim(ob, 1, Sell, 50000, 40)
	trades := ob.Submit(&Order{ID: 2, Side: Buy, Type: Limit, Price: 50000, Quantity: 100, TIF: IOC})
	if len(trades) != 1 || trades[0].Quantity != 40 {
		t.Fatalf("want one 40-qty fill, got %+v", trades)
	}
	if bid, ok := ob.BestBid(); ok {
		t.Errorf("IOC remainder must not rest, but best bid = %d", bid)
	}
}
