package orderbook
import "sort"
type Side string
const (
	Buy  Side = "BUY"
	Sell Side = "SELL"
)
type OrderType string
const (
	Limit  OrderType = "LIMIT"
	Market OrderType = "MARKET"
)
type TIF string
const (
	GTC TIF = "GTC"
	IOC TIF = "IOC"
	FOK TIF = "FOK"
)
type Order struct {
	ID       int64     `json:"sequence_number"`
	Side     Side      `json:"side"`
	Type     OrderType `json:"order_type"`
	Price    int64     `json:"price"`
	Quantity int64     `json:"quantity"`
	TIF      TIF       `json:"tif"`
	seq      uint64
}
type Trade struct {
	MakerID  int64  `json:"maker_id"`
	TakerID  int64  `json:"taker_id"`
	Price    int64  `json:"price"`
	Quantity int64  `json:"quantity"`
	Seq      uint64 `json:"seq"`
}
type level struct {
	price  int64
	orders []*Order
}
func (l *level) totalQty() int64 {
	var s int64
	for _, o := range l.orders {
		s += o.Quantity
	}
	return s
}
type book struct {
	isBid  bool
	levels map[int64]*level
	prices []int64
}
func newBook(isBid bool) *book {
	return &book{isBid: isBid, levels: map[int64]*level{}}
}
func (b *book) best() (int64, bool) {
	if len(b.prices) == 0 {
		return 0, false
	}
	if b.isBid {
		return b.prices[len(b.prices)-1], true
	}
	return b.prices[0], true
}
func (b *book) pricesBestFirst() []int64 {
	out := make([]int64, len(b.prices))
	if b.isBid {
		for i, p := range b.prices {
			out[len(b.prices)-1-i] = p
		}
	} else {
		copy(out, b.prices)
	}
	return out
}
func (b *book) add(o *Order) {
	lvl, ok := b.levels[o.Price]
	if !ok {
		lvl = &level{price: o.Price}
		b.levels[o.Price] = lvl
		i := sort.Search(len(b.prices), func(i int) bool { return b.prices[i] >= o.Price })
		b.prices = append(b.prices, 0)
		copy(b.prices[i+1:], b.prices[i:])
		b.prices[i] = o.Price
	}
	lvl.orders = append(lvl.orders, o)
}
func (b *book) removeLevel(price int64) {
	delete(b.levels, price)
	i := sort.Search(len(b.prices), func(i int) bool { return b.prices[i] >= price })
	if i < len(b.prices) && b.prices[i] == price {
		b.prices = append(b.prices[:i], b.prices[i+1:]...)
	}
}
type OrderBook struct {
	bids     *book
	asks     *book
	byID     map[int64]*Order
	seq      uint64
	tradeSeq uint64
}
func NewOrderBook() *OrderBook {
	return &OrderBook{
		bids: newBook(true),
		asks: newBook(false),
		byID: map[int64]*Order{},
	}
}
func (ob *OrderBook) restingSide(s Side) *book {
	if s == Buy {
		return ob.bids
	}
	return ob.asks
}
func crosses(takerSide Side, takerPrice, restPrice int64) bool {
	if takerSide == Buy {
		return takerPrice >= restPrice
	}
	return takerPrice <= restPrice
}
func (ob *OrderBook) Submit(o *Order) []Trade {
	if o.TIF == "" {
		o.TIF = GTC
	}
	ob.seq++
	o.seq = ob.seq
	opp := ob.asks
	if o.Side == Sell {
		opp = ob.bids
	}
	if o.TIF == FOK && !ob.canFill(o, opp) {
		return nil
	}
	var trades []Trade
	for o.Quantity > 0 {
		bestPrice, ok := opp.best()
		if !ok {
			break
		}
		if o.Type == Limit && !crosses(o.Side, o.Price, bestPrice) {
			break
		}
		lvl := opp.levels[bestPrice]
		maker := lvl.orders[0]
		n := min64(o.Quantity, maker.Quantity)
		ob.tradeSeq++
		trades = append(trades, Trade{
			MakerID:  maker.ID,
			TakerID:  o.ID,
			Price:    bestPrice,
			Quantity: n,
			Seq:      ob.tradeSeq,
		})
		o.Quantity -= n
		maker.Quantity -= n
		if maker.Quantity == 0 {
			lvl.orders = lvl.orders[1:]
			delete(ob.byID, maker.ID)
			if len(lvl.orders) == 0 {
				opp.removeLevel(bestPrice)
			}
		}
	}
	if o.Quantity > 0 && o.Type == Limit && o.TIF == GTC {
		ob.restingSide(o.Side).add(o)
		ob.byID[o.ID] = o
	}
	return trades
}
func (ob *OrderBook) canFill(o *Order, opp *book) bool {
	need := o.Quantity
	for _, p := range opp.pricesBestFirst() {
		if o.Type == Limit && !crosses(o.Side, o.Price, p) {
			break
		}
		need -= opp.levels[p].totalQty()
		if need <= 0 {
			return true
		}
	}
	return false
}
func (ob *OrderBook) Cancel(id int64) bool {
	o, ok := ob.byID[id]
	if !ok {
		return false
	}
	side := ob.restingSide(o.Side)
	lvl := side.levels[o.Price]
	for i, x := range lvl.orders {
		if x.ID == id {
			lvl.orders = append(lvl.orders[:i], lvl.orders[i+1:]...)
			break
		}
	}
	if len(lvl.orders) == 0 {
		side.removeLevel(o.Price)
	}
	delete(ob.byID, id)
	return true
}
func (ob *OrderBook) BestBid() (int64, bool) { return ob.bids.best() }
func (ob *OrderBook) BestAsk() (int64, bool) { return ob.asks.best() }
func (ob *OrderBook) RestingQty(s Side, price int64) int64 {
	if l, ok := ob.restingSide(s).levels[price]; ok {
		return l.totalQty()
	}
	return 0
}
func min64(a, b int64) int64 {
	if a < b {
		return a
	}
	return b
}
