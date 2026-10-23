package outbox

import (
	"encoding/json"
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

// Event is a message waiting to be delivered.
type Event struct {
	ID          string     `gorm:"primaryKey;size:36" json:"id"`
	EventType   string     `gorm:"size:100;not null" json:"eventType"`
	Payload     string     `gorm:"type:text;not null" json:"payload"`
	CreatedAt   time.Time  `json:"createdAt"`
	PublishedAt *time.Time `json:"publishedAt"`
	Attempts    int        `gorm:"not null;default:0" json:"attempts"`
}

// NewEvent creates an event from any JSON-compatible payload.
func NewEvent(eventType string, payload interface{}) (*Event, error) {
	data, err := json.Marshal(payload)
	if err != nil {
		return nil, err
	}

	return &Event{
		ID:        uuid.NewString(),
		EventType: eventType,
		Payload:   string(data),
		CreatedAt: time.Now().UTC(),
	}, nil
}

// Save adds an event using the caller's database transaction.
func Save(tx *gorm.DB, event *Event) error {
	return tx.Create(event).Error
}
