package worker

import (
	"context"
	"fmt"
	"time"

	"day18/logging"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/segmentio/kafka-go"
)

type OutboxRelay struct {
	pool    *pgxpool.Pool
	writer  *kafka.Writer
	logger  *logging.Logger
}

func NewOutboxRelay(pool *pgxpool.Pool, kafkaBroker string, logger *logging.Logger) *OutboxRelay {
	writer := &kafka.Writer{
		Addr:     kafka.TCP(kafkaBroker),
		Topic:    "employee-events",
		Balancer: &kafka.LeastBytes{},
	}
	return &OutboxRelay{pool: pool, writer: writer, logger: logger}
}

func (r *OutboxRelay) Start(ctx context.Context) {
	ticker := time.NewTicker(3 * time.Second)
	defer ticker.Stop()

	for {
		select {
		case <-ctx.Done():
			r.writer.Close()
			return
		case <-ticker.C:
			r.processPendingEvents(ctx)
		}
	}
}

func (r *OutboxRelay) processPendingEvents(ctx context.Context) {
	query := `SELECT id, aggregate_id, event_type, payload FROM outbox_events WHERE status = 'PENDING' LIMIT 10`
	rows, err := r.pool.Query(ctx, query)
	if err != nil {
		return
	}
	defer rows.Close()

	type Event struct {
		ID          int64
		AggregateID int64
		EventType   string
		Payload     []byte
	}

	var events []Event
	for rows.Next() {
		var e Event
		if err := rows.Scan(&e.ID, &e.AggregateID, &e.EventType, &e.Payload); err == nil {
			events = append(events, e)
		}
	}

	for _, e := range events {
		err := r.writer.WriteMessages(ctx, kafka.Message{
			Key:   []byte(fmt.Sprintf("%d", e.AggregateID)),
			Value: e.Payload,
			Headers: []kafka.Header{
				{Key: "event_type", Value: []byte(e.EventType)},
			},
		})

		if err == nil {
			r.pool.Exec(ctx, `UPDATE outbox_events SET status = 'PUBLISHED' WHERE id = $1`, e.ID)
			r.logger.Info("Outbox Relay published event ID: %d to Kafka topic 'employee-events'", e.ID)
		} else {
			r.logger.Error("Outbox Relay failed to publish event ID %d: %v", e.ID, err)
		}
	}
}