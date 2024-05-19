package helpers

import (
	"strconv"

	"github.com/gin-gonic/gin"
)

func GetPagination(c *gin.Context) (int, int) {
	offset := 0
	limit := 1

	if c.Query("offset") != "" {
		parsedOffset, err := strconv.Atoi(c.Query("offset"))
		if err == nil && parsedOffset >= 0 {
			offset = parsedOffset
		}
	}

	if c.Query("limit") != "" {
		parsedLimit, err := strconv.Atoi(c.Query("limit"))
		if err == nil && parsedLimit > 0 {
			limit = parsedLimit
		}
	}

	return offset, limit
}
