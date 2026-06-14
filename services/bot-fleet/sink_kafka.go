package botfleet
import (
	"context"
	"encoding/json"
	"time"
	"github.com/segmentio/kafka-go"
)
type KafkaSink struct{ w *kafka.Writer }
func NewKafkaSink(broker, topic string) *KafkaSink {
	return &KafkaSink{w: &kafka.Writer{
		Addr:                   kafka.TCP(broker),
		Topic:                  topic,
		Balancer:               &kafka.LeastBytes{},
		AllowAutoTopicCreation: true,
		BatchTimeout:           10 * time.Millisecond,
		Async:                  true,
	}}
}
func (s *KafkaSink) Emit(t Telemetry) error {
	b, _ := json.Marshal(t)
	return s.w.WriteMessages(context.Background(), kafka.Message{Key: []byte(t.RunID), Value: b})
}
func (s *KafkaSink) Close() error { return s.w.Close() }
