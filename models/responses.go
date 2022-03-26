package models

// StatusResponse response on status request
type StatusResponse struct {
	Status    bool    `json:"status"`
	Message   string  `json:"message"`
	RequestID string  `json:"requestid"`
	Tables    []Table `json:"tables"`
}

// HeartbeatResponse response on heartbeat request
type HeartbeatResponse struct {
	Status  bool   `json:"status"`
	Message string `json:"message"`
	TableID string `json:"tableid"`
}

// WaiterResponse response on heartbeat request
type WaiterResponse struct {
	Status    bool   `json:"status"`
	Message   string `json:"message"`
	TableID   string `json:"tableid"`
	RequestID string `json:"requestid"`
}
