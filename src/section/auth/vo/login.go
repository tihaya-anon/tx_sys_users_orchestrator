package vo

// LoginRequest
type LoginRequest struct {
	Password string `json:"password"`
	Username string `json:"username"`
}

// LoginResponse
type LoginResponse struct {
	SessionID *string `json:"session_id,omitempty"`
}