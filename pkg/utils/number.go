package utils

import "math"

// GetTotalPage against totalData and limit.
func GetTotalPage(total int64, limit int) int {
	numPages := int(math.Ceil(float64(total)/float64(limit)) + 0.5)
	return numPages
}
