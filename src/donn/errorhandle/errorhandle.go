package errorhandle

import (
	"gopkg.in/gin-gonic/gin.v1"
)

func RespondWithErrorByCustom(code int, message string, c *gin.Context) {
	resp := map[string]string{"status": "_Failure", "error": message}

	c.JSON(code, resp)
	c.Abort()
}

func RespondWithErrorByDefault(errorDefine ErrorDefine, c *gin.Context) {
	//	resp := map[string]string{"status": "_Failure", "error": }
	c.JSON(errorDefine.httpcode, gin.H{
		"status": "_Failure",
		"time":   "time",
		"error":  errorDefine,
	})
	c.Abort()
}
