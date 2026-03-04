package middleware

import (
	"go-auth/internal/auth"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

const (
	ctxUseraIDkey = "auth.userId"
	ctxRolekey    = "auth.role"
)

func AuthRequired(jwtSecret string) gin.HandlerFunc {
	return func (c *gin.Context){
		authHeader := strings.TrimSpace(c.GetHeader("Authorization"))

		if authHeader == "" {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"error":"Missing Authorization token",
			})
			return
		}
		parts := strings.SplitN(authHeader, " ", 2)

		if len(parts) != 2 {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"error":"Invalid Authorization token",
			})
			return
		}
		schema := strings.TrimSpace(parts[0])
		tokenString := strings.TrimSpace(parts[1])

		if !strings.EqualFold(schema, "Bearer"){
			c.AbortWithStatusJSON(http.StatusUnauthorized,gin.H{
				"error":"Authorization Scheme must be Bearer",
			})
			return
		}
		if tokenString == "" {
			c.AbortWithStatusJSON(http.StatusUnauthorized,gin.H{
				"error":"Missing token",
			})
			return
		}
		claims, err := auth.ParseToken(jwtSecret,tokenString) 

		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"error":"Invalid or expired token",
			})
			return
		}
		c.Set(ctxUseraIDkey, claims.Subject)
		c.Set(ctxRolekey, claims.Role)

		c.Next()
	}
}

func GetUserID(c *gin.Context) (string, bool){
	res, ok := c.Get(ctxUseraIDkey)

	if !ok {
		return "", false
	}
	userID, ok := res.(string)
	
	return userID, ok
}

func GetRole(c *gin.Context) (string, bool) {
	res, ok := c.Get(ctxRolekey)

	if !ok {
		return "", false
	}

	role, ok := res.(string)

	return role, ok
}