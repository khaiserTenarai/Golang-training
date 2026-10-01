package model

import "time"

// Department mirrors one row in the "department" table.
type Department struct {
	ID        int
	Name      string
	CreatedAt time.Time
}
