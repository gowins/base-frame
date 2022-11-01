package middleware

import (
	"base-frame/service/bar-single/internal/logic"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Auth(c *gin.Context) {
	if logic.Authentication(c) {
		c.Next()
	} else {
		//如不想在执行其它中间件及handler函数，请调用c.AbortXXX
		c.AbortWithStatus(http.StatusForbidden)
	}
}
