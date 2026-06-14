package botfleet
import (
	"context"
	"math/rand"
	"sync"
	"testing"
)
type fakeEngine struct{}
func (fakeEngine) Send(o Order) (Ack, error) {
	return Ack{SequenceNumber: o.SequenceNumber, Status: "ACK", AckTs: 1}, nil
}
type fakeSink struct {
	mu  sync.Mutex
	got []Telemetry
}
func (s *fakeSink) Emit(t Telemetry) error {
	s.mu.Lock()
	s.got = append(s.got, t)
	s.mu.Unlock()
	return nil
}
func TestRunAuthoritativeRunIDAndUniqueSeq(t *testing.T) {
	sink := &fakeSink{}
	res := Run(context.Background(), fakeEngine{}, sink,
		Config{RunID: "run-X", Bots: 8, OrdersPerBot: 50, Seed: 42})
	if res.Sent != 400 {
		t.Fatalf("want 400 sent, got %d", res.Sent)
	}
	if len(sink.got) != 400 {
		t.Fatalf("want 400 telemetry records, got %d", len(sink.got))
	}
	seen := make(map[int64]bool, 400)
	for _, tm := range sink.got {
		if tm.RunID != "run-X" {
			t.Fatalf("telemetry must carry the authoritative run id, got %q", tm.RunID)
		}
		if seen[tm.SequenceNumber] {
			t.Fatalf("duplicate sequence %d — atomic global counter is broken", tm.SequenceNumber)
		}
		seen[tm.SequenceNumber] = true
	}
	if len(seen) != 400 {
		t.Fatalf("want 400 unique sequences, got %d", len(seen))
	}
}
func TestDeterministicGeneration(t *testing.T) {
	r1 := rand.New(rand.NewSource(7))
	r2 := rand.New(rand.NewSource(7))
	for i := int64(1); i <= 100; i++ {
		if genOrder(r1, i) != genOrder(r2, i) {
			t.Fatalf("genOrder is not deterministic at sequence %d", i)
		}
	}
}
