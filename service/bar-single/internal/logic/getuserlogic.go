package logic

import "github.com/gin-gonic/gin"

func Authentication(c *gin.Context) bool {
	if c.GetHeader("identity") == "owner" {
		return true
	} else {
		return false
	}
}
