package http

import (
	"errors"
	"net/http"
	"strings"

	"github.com/S1nceU/CRMS/apps/api/config"
	"github.com/S1nceU/CRMS/apps/api/domain"
	"github.com/S1nceU/CRMS/apps/api/model/dto"
	_userSer "github.com/S1nceU/CRMS/apps/api/module/user/service"
	"github.com/gin-gonic/gin"
)

type UserHandler struct {
	ser domain.UserService
}

func bindJSON[T any](c *gin.Context) (T, bool) {
	var request T
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"Message": err.Error(),
		})
		return request, false
	}
	return request, true
}

func NewUserHandler(r gin.IRoutes, ser domain.UserService) {
	handler := &UserHandler{
		ser: ser,
	}

	r.POST("/userLogin", handler.Login)
	r.POST("/userAuthentication", handler.Authentication)
	r.POST("/userLogout", handler.Logout)

}

// Login @Summary Login
// @Description Login
// @Tags User
// @Accept json
// @Produce application/json
// @Param UserLoginRequest body model.UserLoginRequest true "User Login Request"
// @Success 200 {object} string
// @Router /userLogin [post]
func (u *UserHandler) Login(c *gin.Context) {
	request, ok := bindJSON[dto.UserLoginRequest](c)
	if !ok {
		return
	}
	token, err := u.ser.Login(request.Username, request.Password)

	if err != nil {
		if err.Error() == "user not found" {
			c.JSON(http.StatusOK, gin.H{
				"Message": err.Error(),
			})
			return
		}
		if err.Error() == "password is incorrect" {
			c.JSON(http.StatusOK, gin.H{
				"Message": err.Error(),
			})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{
			"Message": err.Error(),
		})
		return
	}
	// Set cookie with SameSite and Secure according to config
	cookie := &http.Cookie{
		Name:     "token",
		Value:    token,
		Path:     "/",
		MaxAge:   int(_userSer.TokenExpireDuration().Seconds()),
		HttpOnly: true,
		Secure:   config.Val.CookieSecure,
	}
	// Only set SameSite=None when using secure cookies (required by browsers)
	if config.Val.CookieSecure {
		cookie.SameSite = http.SameSiteNoneMode
	}
	http.SetCookie(c.Writer, cookie) // When CRMS runs in the docker container, the domain should be changed to "localhost"
	c.JSON(http.StatusOK, gin.H{
		"Message": "Login successfully",
		"token":   token,
	})
}

// Authentication @Summary Authentication
// @Description Authentication
// @Tags User
// @Accept json
// @Produce application/json
// @Param UserTokenRequest body model.UserTokenRequest true "User JWT Token"
// @Success 200 {object} string
// @Router /userAuthentication [post]
func (u *UserHandler) Authentication(c *gin.Context) {
	_, ok := bindJSON[dto.UserTokenRequest](c)
	if !ok {
		return
	}
	token, err := extractToken(c)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"Message": "Authentication failed",
		})
		return
	}

	username, err := u.ser.Authentication(token)

	if err != nil {
		if strings.HasPrefix(err.Error(), "token is expired") {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"Message": err.Error(),
			})
			return
		}
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"Message": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"Message":  "Authentication successfully",
		"username": username,
	})
}

// Logout @Summary Logout
// @Description Logout
// @Tags User
// @Accept json
// @Produce application/json
// @Param UserTokenRequest body model.UserTokenRequest true "User JWT Token"
// @Success 200 {object} string
// @Router /userLogout [post]
func (u *UserHandler) Logout(c *gin.Context) {
	_, ok := bindJSON[dto.UserTokenRequest](c)
	if !ok {
		return
	}

	token, err := extractToken(c)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"Message": "Authentication failed",
		})
		return
	}

	_, err = u.ser.Authentication(token)

	if err != nil {
		if strings.HasPrefix(err.Error(), "token is expired") {
			c.JSON(http.StatusUnauthorized, gin.H{
				"Message": err.Error(),
			})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{
			"Message": err.Error(),
		})
		return
	}
	// Clear cookie
	cookie := &http.Cookie{
		Name:     "token",
		Value:    "",
		Path:     "/",
		MaxAge:   -1,
		HttpOnly: true,
		Secure:   config.Val.CookieSecure,
	}
	if config.Val.CookieSecure {
		cookie.SameSite = http.SameSiteNoneMode
	}
	http.SetCookie(c.Writer, cookie) // When CRMS runs in the docker container, the domain should be changed to "localhost"
	c.JSON(http.StatusOK, gin.H{
		"Message": "Logout successfully",
	})
}

func extractToken(c *gin.Context) (string, error) {

	if cookie, err := c.Cookie("token"); err == nil && cookie != "" {
		return cookie, nil
	}
	authHeader := c.GetHeader("Authorization")
	if strings.HasPrefix(authHeader, "Bearer ") {
		token := strings.TrimPrefix(authHeader, "Bearer ")
		if token != "" {
			return token, nil
		}
	}
	return "", errors.New("missing token")
}
