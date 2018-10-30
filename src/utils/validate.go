package utils

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// ValidateRequestForm s
func ValidateRequestForm(c *gin.Context, m map[string]string) bool {
	// keys := make([]int, 0, len(m))
	for k := range m {
		// keys = append(keys, k)
		v := c.PostForm(k)
		if v == "" {
			c.JSON(http.StatusOK, StandardEmptyMessage(k))
			return false
		}
		m[k] = v
	}
	return true
}
