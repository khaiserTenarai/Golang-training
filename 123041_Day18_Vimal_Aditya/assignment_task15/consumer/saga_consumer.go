package consumer

import (
	"context"
	"encoding/json"

	"day18/logging"

	"day18/model"
	"day18/repository"

	"github.com/segmentio/kafka-go"
)

type SagaConsumer struct {
	reader *kafka.Reader
	repo   repository.EmployeeRepository
	logger *logging.Logger
}

func NewSagaConsumer(kafkaBroker string, repo repository.EmployeeRepository, logger *logging.Logger) *SagaConsumer {
	reader := kafka.NewReader(kafka.ReaderConfig{
		Brokers:  []string{kafkaBroker},
		Topic:    "department-events",
		GroupID:  "employee-saga-group",
		MinBytes: 10,
		MaxBytes: 10e6,
	})
	return &SagaConsumer{reader: reader, repo: repo, logger: logger}
}

func (c *SagaConsumer) Start(ctx context.Context) {
	go func() {
		for {
			select {
			case <-ctx.Done():
				c.reader.Close()
				return
			default:
				msg, err := c.reader.ReadMessage(ctx)
				if err != nil {
					continue
				}

				var event struct {
					EmployeeID int64  `json:"employee_id"`
					Reason     string `json:"reason"`
				}

				if err := json.Unmarshal(msg.Value, &event); err == nil {
					for _, header := range msg.Headers {
						if header.Key == "event_type" && string(header.Value) == "DepartmentReservationFailed" {
							// Compensating Transaction: Rollback employee state to REJECTED
							c.repo.UpdateStatus(ctx, event.EmployeeID, model.StatusRejected)
							c.logger.Info("Saga Rollback executed: Employee ID %d rejected due to: %s", event.EmployeeID, event.Reason)
						}
					}
				}
			}
		}
	}()
}