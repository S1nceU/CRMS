package route

import (
	"net/http"
	"strings"

	"github.com/S1nceU/CRMS/apps/api/domain"
	"github.com/gin-gonic/gin"
)

func AuthMiddleware(userSvc domain.UserService) gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenString := ""
		if cookie, err := c.Cookie("token"); err == nil && cookie != "" {
			tokenString = cookie
		}
		if tokenString == "" {
			authHeader := c.GetHeader("Authorization")
			if strings.HasPrefix(authHeader, "Bearer ") {
				tokenString = strings.TrimPrefix(authHeader, "Bearer ")
			}
		}
		if tokenString == "" {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"Message": "Unauthorized"})
			return
		}

		username, err := userSvc.Authentication(tokenString)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"Message": err.Error()})
			return
		}
		c.Set("username", username)
		c.Next()
	}
}
