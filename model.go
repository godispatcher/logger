package logger

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"time"
)

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

func NewLoggedRequest(r *http.Request) (LoggedRequest, error) {
	bodyBytes, err := io.ReadAll(r.Body)
	if err != nil {
		return LoggedRequest{}, err
	}

	r.Body = io.NopCloser(bytes.NewBuffer(bodyBytes))

	return LoggedRequest{
		Method:  r.Method,
		URL:     r.URL.String(),
		Headers: r.Header,
		Body:    json.RawMessage(bodyBytes),
	}, nil
}

func NewLoggedResponse(statusCode int, headers http.Header, body interface{}) LoggedResponse {
	return LoggedResponse{
		StatusCode: statusCode,
		Headers:    headers,
		Body:       body,
	}
}
