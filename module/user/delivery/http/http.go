package http

import (
	"github.com/S1nceU/CRMS/domain"
	"github.com/S1nceU/CRMS/model"
	_userSer "github.com/S1nceU/CRMS/module/user/service"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
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
	request := model.UserLoginRequest{}
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"msg": err.Error()})
		return
	}
	token, err := u.ser.Login(request.Username, request.Password)

	if err != nil {
		if err.Error() == "user not found" {
			c.JSON(http.StatusOK, gin.H{"msg": err.Error()})
			return
		}
		if err.Error() == "password is incorrect" {
			c.JSON(http.StatusOK, gin.H{"msg": err.Error()})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"msg": err.Error()})
		return
	}
	c.SetCookie("token", token, int(_userSer.TokenExpireDuration.Seconds()), "/", "localhost:8090", false, true) // When the docker container is running, the domain should be changed to "localhost"
	c.JSON(http.StatusOK, gin.H{
		"msg":   "Login successfully",
		"token": token,
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
	request := model.UserTokenRequest{}
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if _, err := c.Cookie("token"); err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"msg": "Authentication failed"})
		return
	}

	username, err := u.ser.Authentication(request.Token)

	if err != nil {
		if strings.HasPrefix(err.Error(), "token is expired") {
			c.JSON(http.StatusUnauthorized, gin.H{"msg": err.Error()})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"msg": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"msg":      "Authentication successfully",
		"username": username,
	})
}
