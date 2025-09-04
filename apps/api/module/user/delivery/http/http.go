package http

import (
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

func NewUserHandler(e *gin.Engine, ser domain.UserService) {
	handler := &UserHandler{
		ser: ser,
	}
	api := e.Group("/api")
	{
		api.POST("/userLogin", handler.Login)
		api.POST("/userAuthentication", handler.Authentication)
		api.POST("/userLogout", handler.Logout)
	}
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
	request := dto.UserLoginRequest{}
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"Message": err.Error(),
		})
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
		MaxAge:   int(_userSer.TokenExpireDuration.Seconds()),
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
	request := dto.UserTokenRequest{}
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"Message": err.Error(),
		})
		return
	}

	if _, err := c.Cookie("token"); err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{
			"Message": "Authentication failed",
		})
		return
	}

	username, err := u.ser.Authentication(request.Token)

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
	request := dto.UserTokenRequest{}
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"Message": err.Error(),
		})
		return
	}

	if _, err := c.Cookie("token"); err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{
			"Message": "Not logged in yet",
		})
		return
	}

	_, err := u.ser.Authentication(request.Token)

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
