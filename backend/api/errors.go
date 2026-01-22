package api

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

type FieldError struct {
	Field   string `json:"field"`
	Message string `json:"message"`
}

type ErrorResponse struct {
	Status    int          `json:"status"`
	Error     string       `json:"error"`
	Message   string       `json:"message"`
	Details   []FieldError `json:"details,omitempty"`
	Timestamp string       `json:"timestamp"`
}

func SendError(c *gin.Context, status int, message string, details ...FieldError) {
	c.JSON(status, ErrorResponse{
		Status:    status,
		Error:     http.StatusText(status),
		Message:   message,
		Details:   details,
		Timestamp: time.Now().Format(time.RFC3339),
	})
}
