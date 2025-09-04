package route

import (
    "github.com/S1nceU/CRMS/apps/api/config"
    "github.com/gin-gonic/gin"
    "net/http"
)

func Cors() gin.HandlerFunc {
    return func(c *gin.Context) {
        method := c.Request.Method
        origin := c.Request.Header.Get("Origin")

        // Only echo back a specific allowed origin to enable credentials
        if origin != "" && (config.Val.FrontendOrigin == "*" || origin == config.Val.FrontendOrigin) {
            c.Header("Access-Control-Allow-Origin", origin)
            c.Header("Vary", "Origin")
        }

        c.Header("Access-Control-Allow-Headers", "Content-Type,AccessToken,X-CSRF-Token, Authorization, Token, Application, Access-Control-Allow-Origin")
        c.Header("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
        c.Header("Access-Control-Expose-Headers", "Content-Length, Access-Control-Allow-Origin, Access-Control-Allow-Headers, Content-Type")
        c.Header("Access-Control-Allow-Credentials", "true")

        if method == "OPTIONS" {
            c.AbortWithStatus(http.StatusNoContent)
            return
        }

        c.Next()
    }
}
