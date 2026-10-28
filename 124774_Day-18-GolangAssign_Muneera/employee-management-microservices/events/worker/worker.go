package worker

import (
	"bytes"
	"context"
	"encoding/json"
	"log"
	"net/http"
	"time"

	"employee-management-microservices/events/outbox"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

// Worker sends pending outbox events to another service.
type Worker struct {
	db       *gorm.DB
	target   string
	client   *http.Client
	interval time.Duration
}

func New(db *gorm.DB, target string) *Worker {
	return &Worker{
		db:       db,
		target:   target,
		client:   &http.Client{Timeout: 5 * time.Second},
		interval: 2 * time.Second,
	}
}

// Run keeps checking for pending events until the context is cancelled.
func (w *Worker) Run(ctx context.Context) {
	ticker := time.NewTicker(w.interval)
	defer ticker.Stop()

	for {
		if err := w.ProcessBatch(ctx); err != nil {
			log.Printf("Outbox worker error: %v", err)
		}

		select {
		case <-ctx.Done():
			return
		case <-ticker.C:
		}
	}
}

// ProcessBatch claims and sends a small batch of events.
func (w *Worker) ProcessBatch(ctx context.Context) error {
	var events []outbox.Event

	// Claim pending rows in a short transaction.
	err := w.db.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		return tx.Clauses(clause.Locking{
			Strength: "UPDATE",
			Options:  "SKIP LOCKED",
		}).
			Where("published_at IS NULL").
			Order("created_at").
			Limit(20).
			Find(&events).Error
	})
	if err != nil {
		return err
	}

	for _, event := range events {
		if ctx.Err() != nil {
			return ctx.Err()
		}

		// Send the event ID, type, and original payload.
		body, err := json.Marshal(map[string]interface{}{
			"eventId":   event.ID,
			"eventType": event.EventType,
			"payload":   json.RawMessage(event.Payload),
		})
		if err != nil {
			log.Printf("Invalid outbox event %s: %v", event.ID, err)
			continue
		}

		req, err := http.NewRequestWithContext(
			ctx, http.MethodPost, w.target, bytes.NewReader(body),
		)
		if err != nil {
			return err
		}
		req.Header.Set("Content-Type", "application/json")

		resp, err := w.client.Do(req)
		if err != nil {
			_ = w.db.Model(&event).UpdateColumn(
				"attempts", gorm.Expr("attempts + ?", 1),
			).Error
			log.Printf("Event %s delivery failed: %v", event.ID, err)
			continue
		}
		_ = resp.Body.Close()

		if resp.StatusCode < 200 || resp.StatusCode >= 300 {
			_ = w.db.Model(&event).UpdateColumn(
				"attempts", gorm.Expr("attempts + ?", 1),
			).Error
			log.Printf("Event %s returned HTTP %d", event.ID, resp.StatusCode)
			continue
		}

		now := time.Now().UTC()
		if err := w.db.Model(&event).Updates(map[string]interface{}{
			"published_at": now,
			"attempts":     gorm.Expr("attempts + ?", 1),
		}).Error; err != nil {
			log.Printf("Could not mark event %s as published: %v", event.ID, err)
		}
	}

	return nil
}
