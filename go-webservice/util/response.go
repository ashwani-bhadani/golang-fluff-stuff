package util

import "github.com/gin-gonic/gin"

func HandleError(c *gin.Context, status int, message string) {
    c.AbortWithStatusJSON(status, gin.H{
        "error":   true,
        "message": message,
    })
}