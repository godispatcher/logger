package logger

import "time"

type LogEntry struct {
	Timestamp    time.Time      `json:"timestamp"`
	Request      LoggedRequest  `json:"request"`
	Response     LoggedResponse `json:"response"`
	Duration     time.Duration  `json:"duration"`
	ErrorMessage string         `json:"error_message,omitempty"`
}
type LoggedRequest struct {
	Method  string              `json:"method"`
	URL     string              `json:"url"`
	Headers map[string][]string `json:"headers"`
	Body    interface{}         `json:"body"`
}
type LoggedResponse struct {
	StatusCode int                 `json:"status_code"`
	Headers    map[string][]string `json:"headers"`
	Body       interface{}         `json:"body"`
}
