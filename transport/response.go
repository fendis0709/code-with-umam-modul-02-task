package transport

// StatusResponse represents a standard status response.
type StatusResponse struct {
	Code   int    `json:"code"`
	Status string `json:"status"`
}
