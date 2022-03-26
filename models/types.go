package models

import "time"

// Table represent physical table in bar
type Table struct {
	TableID         string       `json:"table_id"`
	LastHeartbeatAt time.Time    `json:"last_heartbeat_at"`
	Events          []TableEvent `json:"events"`
}

// TableEvent stores events from table - waiter/bill requests, etc
type TableEvent struct {
	EventID   string    `json:"event_id"`
	EventType string    `json:"event_type"`
	CreatedAt time.Time `json:"created_at"`
	IsHandled bool      `json:"is_handled"`
}
