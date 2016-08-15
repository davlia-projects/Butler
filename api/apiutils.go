package api

import (
	"encoding/json"
	"net/http"
	"strconv"
)

type ErrorResponse struct {
	Status     int    `json:"status"`
	Message    string `json:"message"`
	StatusText string `json:"error"`
}

func NewErrorResponse(status int, message string) ErrorResponse {
	statusText := http.StatusText(status)
	if statusText == "" {
		statusText = ExtentionStatusText(status)
	}
	return ErrorResponse{
		Status:     status,
		Message:    message,
		StatusText: statusText,
	}
}

func ServeJSON(w http.ResponseWriter, v interface{}) {
	content, err := json.MarshalIndent(v, "", "  ")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Length", strconv.Itoa(len(content)))
	w.Header().Set("Content-Type", "application/json")
	w.Write(content)
}

func ServeError(w http.ResponseWriter, errRes ErrorResponse) {
	w.WriteHeader(errRes.Status)
	ServeJSON(w, errRes)
}
