package util

import "github.com/WayeeeX/go-gin-example/models/request"

// GetPage get page parameters
func GetOffset(req request.PageQuery) (offset int) {
	return (req.PageNum - 1) * req.PageSize
}
