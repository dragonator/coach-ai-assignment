package consumer

import (
	"context"
	"fmt"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/push"
	io_prometheus_client "github.com/prometheus/client_model/go"

	"github.com/dragonator/coach-ai-assignment/internal/config"
	kafkainternal "github.com/dragonator/coach-ai-assignment/internal/kafka"
)

var (
	// Define Prometheus metric
	_kafkaMessagesConsumedTotal = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Name: "kafka_messages_consumed_total",
			Help: "Total number of Kafka messages consumed",
		},
		[]string{"topic"},
	)
)

func init() {
	prometheus.MustRegister(_kafkaMessagesConsumedTotal)
}

type MetricsDecorator struct {
	metricsPusher *push.Pusher
	next          *Service
	topic         string
	instanceID    string
}

func NewMetricsDecorator(
	config *config.Config,
	metricsPusher *push.Pusher,
	next *Service,
	topic string,
) *MetricsDecorator {
	return &MetricsDecorator{
		metricsPusher: metricsPusher,
		next:          next,
		topic:         topic,
		instanceID:    config.InstanceID,
	}
}

func (d *MetricsDecorator) ProcessEvent(ctx context.Context, event kafkainternal.Event) error {
	if err := d.next.ProcessEvent(ctx, event); err != nil {
		return fmt.Errorf("calling next: %w", err)
	}

	_kafkaMessagesConsumedTotal.WithLabelValues(d.topic).Inc()

	err := d.metricsPusher.Grouping("instance", d.instanceID).Add()
	if err != nil {
		return fmt.Errorf("pushing metrics: %w", err)
	}

	m := &io_prometheus_client.Metric{}
	if err := _kafkaMessagesConsumedTotal.WithLabelValues(d.topic).Write(m); err == nil {
		fmt.Printf("Current counter value: %v\n", m.GetCounter().GetValue())
	} else {
		return fmt.Errorf("checking metric: %w", err)
	}

	return nil
}
