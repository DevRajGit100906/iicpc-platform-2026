package scoring
import "testing"
func TestCorrectnessGateOutranksSpeed(t *testing.T) {
	cfg := DefaultConfig()
	fastCorrect := Compute(Input{Lat: Percentiles{P99: 500}, TPS: 8000, Correctness: 1, CorrectnessPassed: true}, cfg)
	slowCorrect := Compute(Input{Lat: Percentiles{P99: 4000}, TPS: 2000, Correctness: 1, CorrectnessPassed: true}, cfg)
	fastWrong := Compute(Input{Lat: Percentiles{P99: 500}, TPS: 8000, Correctness: 0.8, CorrectnessPassed: false}, cfg)
	if !(fastCorrect.Composite > slowCorrect.Composite) {
		t.Fatalf("fast+correct (%.1f) should beat slow+correct (%.1f)", fastCorrect.Composite, slowCorrect.Composite)
	}
	if !(slowCorrect.Composite > fastWrong.Composite) {
		t.Fatalf("GATE FAILED: slow+correct (%.1f) must beat fast+wrong (%.1f)", slowCorrect.Composite, fastWrong.Composite)
	}
	if !fastWrong.Gated {
		t.Fatal("an engine that failed correctness must be marked gated")
	}
	t.Logf("fast+correct=%.1f  slow+correct=%.1f  fast+wrong=%.1f(gated)",
		fastCorrect.Composite, slowCorrect.Composite, fastWrong.Composite)
}
func TestPerfectEngineScores100(t *testing.T) {
	r := Compute(Input{Lat: Percentiles{P99: 200}, TPS: 9000, Correctness: 1, CorrectnessPassed: true}, DefaultConfig())
	if r.Composite < 99.999 {
		t.Fatalf("an engine beating every target should score 100, got %.4f", r.Composite)
	}
}
func TestPercentiles(t *testing.T) {
	var us []int64
	for i := int64(1); i <= 100; i++ {
		us = append(us, i)
	}
	p := ComputePercentiles(us)
	if p.Max != 100 {
		t.Fatalf("max: want 100, got %d", p.Max)
	}
	if p.P50 < 49 || p.P50 > 52 {
		t.Fatalf("p50 out of range: %d", p.P50)
	}
	if p.P99 < 98 {
		t.Fatalf("p99 too low: %d", p.P99)
	}
}
