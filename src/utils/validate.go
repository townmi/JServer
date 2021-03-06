package utils

import (
	"errors"

	"github.com/gin-gonic/gin"
)

// ValidateRequestForm s
func ValidateRequestForm(c *gin.Context, m map[string]string) error {
	// keys := make([]int, 0, len(m))
	for k := range m {
		// keys = append(keys, k)
		v := c.PostForm(k)
		if v == "" {
			return errors.New(StandardEmptyMessage(k))
		}
		m[k] = v
	}
	return nil
}
