package middleware

import (
	"donn/errorhandle"
	"fmt"
	"gopkg.in/gin-gonic/gin.v1"
	"net/http"
)

func Auth0Middleware() gin.HandlerFunc {

	fmt.Println("Auth0 Middleware Loading...")

	return func(c *gin.Context) {
		fmt.Println("Check Auth")

		// before request
		token := c.Request.Header.Get("token")
		fmt.Println(token)

		if token == "" {
			//			errorhandle.RespondWithError(http.StatusUnauthorized, "API token required", c)
			errorhandle.RespondWithErrorByDefault(errorhandle.Err101, c)
			return
		}

		if token != "3345678" {
			errorhandle.RespondWithErrorByCustom(http.StatusUnauthorized, "Invalid API token", c)
			return
		}

		c.Next()

		// after request
	}
}
