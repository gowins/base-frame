package middleware

import (
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
)

func Auth(c *gin.Context) {
	demoValue := time.Now().String() + " authMiddle"
	fmt.Printf("this auth middleware\n")
	c.Set("authMiddle", demoValue)
	c.Next()
}
