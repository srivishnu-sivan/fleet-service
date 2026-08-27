package auth

import (
	"fmt"
	"strings"

	"github.com/gin-gonic/gin"
)

func (j *JWTService) AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// we'll implement this
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			c.JSON(401, gin.H{"error": "missing authorization header"})
			c.Abort()
			return
		}

		parts := strings.SplitN(authHeader, " ", 2)
		fmt.Println(parts)
		fmt.Println("AUTH HEADER:", authHeader)
fmt.Println("TOKEN:", parts[1])

		if len(parts) != 2 || parts[0] != "Bearer" || parts[1] == "" {
			c.JSON(401, gin.H{"error": "invalid authorization header"})
			c.Abort()
			return
		}

		claims, err := j.VerifyToken(parts[1])
if err != nil {
	c.JSON(401, gin.H{"error": "invalid or expired token"})
	c.Abort()
	return
}

c.Set("claims", claims)
c.Next()
	}
}
