package http

import (
	"github.com/S1nceU/CRMS/domain"
	"github.com/S1nceU/CRMS/model"
	"github.com/gin-gonic/gin"
	"net/http"
)

type CitizenshipHandler struct {
	service domain.CitizenshipService
}

func NewCitizenshipHandler(e *gin.Engine, service domain.CitizenshipService) {
	handler := &CitizenshipHandler{
		service: service,
	}
	api := e.Group("/api")
	{
		api.POST("/citizenships", handler.ListCitizenships)
		api.POST("/citizenshipId", handler.GetCitizenshipByID)
		api.POST("/citizenshipNation", handler.GetCitizenshipByCitizenshipName)
	}
}

// ListCitizenships @Summary ListCitizenships
// @Description List all citizenships
// @Tags Citizenship
// @Accept json
// @Produce application/json
// @Success 200 {object} []model.Citizenship
// @Router /citizenships [post]
func (u *CitizenshipHandler) ListCitizenships(c *gin.Context) {
	citizenships, err := u.service.ListCitizenships()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"citizenships": citizenships})
}

// GetCitizenshipByID @Summary GetCitizenshipByID
// @Description Get citizenship by ID
// @Tags Citizenship
// @Accept json
// @Produce application/json
// @Param CitizenshipId body model.CitizenshipRequest true "Citizenship ID"
// @Success 200 {object} model.Citizenship
// @Router /citizenshipId [post]
func (u *CitizenshipHandler) GetCitizenshipByID(c *gin.Context) {
	request := model.CitizenshipRequest{}
	if err := c.BindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"Message": err.Error(),
		})
		return
	}
	citizenship, err := u.service.GetCitizenshipByID(request.CitizenshipId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"citizenship": citizenship})
}

// GetCitizenshipByCitizenshipName @Summary GetCitizenshipByCitizenshipName
// @Description Get citizenship by CitizenshipName
// @Tags Citizenship
// @Accept json
// @Produce application/json
// @Param CitizenshipNation body model.CitizenshipNameRequest true "Citizenship Nation"
// @Success 200 {object} model.Citizenship
// @Router /citizenshipNation [post]
func (u *CitizenshipHandler) GetCitizenshipByCitizenshipName(c *gin.Context) {
	request := model.CitizenshipNameRequest{}
	if err := c.BindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"Message": err.Error(),
		})
		return
	}
	citizenship, err := u.service.GetCitizenshipByCitizenshipName(request.CitizenshipName)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"citizenship": citizenship})
}
