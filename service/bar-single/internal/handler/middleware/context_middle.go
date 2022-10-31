package middleware

import (
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
)

func ContextM(c *gin.Context) {
	demoValue := time.Now().String() + " contextMiddle"
	fmt.Printf("this auth context middleware\n")
	c.Set("contextMiddle", demoValue)
	c.Next()
}
