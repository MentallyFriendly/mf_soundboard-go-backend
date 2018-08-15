package utils

// Result type
type Result struct {
	Success *Success `json:"success"`
	Error   *Error   `json:"error"`
}

// Success type
type Success struct {
	StatusCode int         `json:"status_code"`
	Data       interface{} `json:"data"`
}

// Error type
type Error struct {
	StatusCode int    `json:"status_code"`
	StatusText string `json:"error_text"`
}
