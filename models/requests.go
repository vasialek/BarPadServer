package models

// StatusRequest status request
type StatusRequest struct {
	RequestID string `json:"requestid"`
}

// HeartbeatRequest heartbeat request
type HeartbeatRequest struct {
	TableID string `json:"tableid"`
}

// WaiterRequest waiter request
type WaiterRequest struct {
	TableID   string `json:"tableid"`
	RequestID string `json:"requestid"`
}
