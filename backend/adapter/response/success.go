package response

import (
	"encoding/json"
	"net/http"
)

type Success struct {
	statusCode int
	result     any
}

func NewSuccess(result any, status int) *Success {
	return &Success{
		statusCode: status,
		result:     result,
	}
}

func (s *Success) Send(w http.ResponseWriter) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(s.statusCode)
	return json.NewEncoder(w).Encode(s.result)
}
