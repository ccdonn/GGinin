package middleware

import (
	"fmt"
	"gopkg.in/gin-gonic/gin.v1"
	"time"
)

func DummyMiddle() gin.HandlerFunc {

	fmt.Println("Dummy Middleware Loading...")

	return func(c *gin.Context) {
		// before request
		// fmt.Println("inDummy")
		t := time.Now()

		c.Next()

		// after request
		latency := time.Since(t)
		fmt.Println(latency)
	}
}
