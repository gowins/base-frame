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

func ContextM2(c *gin.Context) {
	demoValue := time.Now().String() + " contextMiddle2"
	fmt.Printf("this auth context middleware\n")
	c.Set("contextMiddle2", demoValue)
	c.Next()
}
