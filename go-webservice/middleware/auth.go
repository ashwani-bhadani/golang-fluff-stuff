package middleware

import (
    "github.com/gin-gonic/gin"
    "net/http"
)

func AuthMiddleware() gin.HandlerFunc {
    return func(c *gin.Context) {
        // Simulated auth for demo
        token := c.GetHeader("Authorization")
        if token == "" {
            c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
            return
        }
        c.Set("role", "admin") // Simulated role
        c.Next()
    }
}

func AdminOnly() gin.HandlerFunc {
    return func(c *gin.Context) {
        role := c.GetString("role")
        if role != "admin" {
            c.AbortWithStatusJSON(http.StatusForbidden, gin.H{"error": "forbidden"})
            return
        }
        c.Next()
    }
}