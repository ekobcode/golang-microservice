package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

// TransactionIDMiddleware adds or generates a transaction ID per request
func TransactionIDMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		transactionID := c.GetHeader("X-Transaction-ID")
		if transactionID == "" {
			transactionID = uuid.New().String()
		}

		// Store it in Gin context
		c.Set("transaction_id", transactionID)

		// Include it in response header for traceability
		c.Writer.Header().Set("X-Transaction-ID", transactionID)

		c.Next()
	}
}
