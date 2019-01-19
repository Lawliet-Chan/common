package errhandle

import (
	"github.com/CrocdileChan/common/logger"
	"github.com/gin-gonic/gin"
)

func BadRequestCheck(c *gin.Context, err error, msg string) {
	HttpCheck(c, 400, err, msg)
}

func ForbiddenCheck(c *gin.Context, err error, msg string) {
	HttpCheck(c, 403, err, msg)
}

func InternalErrorCheck(c *gin.Context, err error, msg string) {
	HttpCheck(c, 500, err, msg)
}

func UnauthCheck(c *gin.Context, err error, msg string) {
	HttpCheck(c, 401, err, msg)
}

func HttpCheck(c *gin.Context, httpcode int, err error, msg string) {
	if err != nil {
		msg += ": %v"
		logger.GetLogger().Errorf(msg, err)
		c.AbortWithStatus(httpcode)
	}
}
