package utils

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type ErrorResponse struct {
	Status        string `json:"status"`
	Message       string `json:"message"`
	TransactionID string `json:"transaction_id,omitempty"`
}

func RespondError(c *gin.Context, logger *zap.SugaredLogger, err error, message string, status int) {
	transactionID, _ := c.Get("transaction_id")

	// Structured error log with transaction_id
	logger.Errorw("API error",
		"transaction_id", transactionID,
		"endpoint", c.FullPath(),
		"method", c.Request.Method,
		"status", status,
		"error", err,
	)

	c.JSON(status, ErrorResponse{
		Status:        "error",
		Message:       message,
		TransactionID: transactionID.(string),
	})
	c.Abort()
}
