package util

import (
	"github.com/gin-gonic/gin"
)

func BindJSONWithError(c *gin.Context, code int, err error) {
	c.AbortWithStatusJSON(code, gin.H{"error": err.Error()})
}

func BindJSONWithMessage(c *gin.Context, code int, message string) {
	c.AbortWithStatusJSON(code, gin.H{"message": message})
}