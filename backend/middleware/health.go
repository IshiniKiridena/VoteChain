package middleware

import (
    "github.com/gin-gonic/gin"
)

func Health(c *gin.Context) {
    c.JSON(200, gin.H{"message": "VoteChain server is up and running on port 8000"})
    c.Next()
}
