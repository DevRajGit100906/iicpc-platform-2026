package scoring
import "sort"
type Percentiles struct {
	P50 int64
	P90 int64
	P99 int64
	Max int64
}
func ComputePercentiles(us []int64) Percentiles {
	if len(us) == 0 {
		return Percentiles{}
	}
	s := append([]int64(nil), us...)
	sort.Slice(s, func(i, j int) bool { return s[i] < s[j] })
	pick := func(q float64) int64 {
		idx := int(q * float64(len(s)-1))
		if idx < 0 {
			idx = 0
		}
		if idx >= len(s) {
			idx = len(s) - 1
		}
		return s[idx]
	}
	return Percentiles{P50: pick(0.50), P90: pick(0.90), P99: pick(0.99), Max: s[len(s)-1]}
}
type Weights struct {
	Lat float64
	TPS float64
	Cor float64
}
type Config struct {
	TargetP99US int64
	TargetTPS   float64
	Weights     Weights
	GateCap     float64
}
func DefaultConfig() Config {
	return Config{
		TargetP99US: 1000,
		TargetTPS:   5000,
		Weights:     Weights{Lat: 0.30, TPS: 0.25, Cor: 0.45},
		GateCap:     0.25,
	}
}
type Input struct {
	Lat               Percentiles
	TPS               float64
	Correctness       float64
	CorrectnessPassed bool
}
type Result struct {
	Composite float64
	LatScore  float64
	TPSScore  float64
	CorScore  float64
	Gated     bool
	P99US     int64
	TPS       float64
	Correct   float64
}
func clamp01(x float64) float64 {
	if x < 0 {
		return 0
	}
	if x > 1 {
		return 1
	}
	return x
}
func Compute(in Input, cfg Config) Result {
	lat := 1.0
	if in.Lat.P99 > 0 {
		lat = clamp01(float64(cfg.TargetP99US) / float64(in.Lat.P99))
	}
	tps := clamp01(in.TPS / cfg.TargetTPS)
	cor := clamp01(in.Correctness)
	w := cfg.Weights
	composite := (w.Lat*lat + w.TPS*tps + w.Cor*cor) * 100.0
	gated := false
	if !in.CorrectnessPassed {
		if cap := cfg.GateCap * 100.0; composite > cap {
			composite = cap
			gated = true
		}
	}
	return Result{
		Composite: composite,
		LatScore:  lat, TPSScore: tps, CorScore: cor,
		Gated: gated,
		P99US: in.Lat.P99, TPS: in.TPS, Correct: cor,
	}
}
