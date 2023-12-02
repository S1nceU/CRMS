package http

import (
	"github.com/S1nceU/CRMS/model"
	"github.com/S1nceU/CRMS/module/history"
	"github.com/gin-gonic/gin"
	"strconv"
	"time"
)

type HistoryHandler struct {
	ser history.Service
}

func NewHistoryHandler(e *gin.Engine, ser history.Service) {
	handler := &HistoryHandler{
		ser: ser,
	}
	e.GET("/api/historyList", handler.ListHistories)
	e.GET("/api/history", handler.GetHistory)
	e.POST("/api/history", handler.CreateHistory)
	e.PUT("/api/history", handler.ModifyHistory)
	e.DELETE("/api/history", handler.DeleteHistory)
}

// ListHistories @Summary ListHistories
// @Description Get all History
// @Accept json
// @Tags History
// @Produce application/json
// @Success 200 {object} model.History
// @Failure 500 {string} string "{"Message": "Internal Error!"}"
// @Router /historyList [get]
func (u *HistoryHandler) ListHistories(c *gin.Context) {
	historyList, err := u.ser.ListHistories()
	if err != nil {
		c.JSON(500, gin.H{
			"Message": "Internal Error!",
		})
		return
	}
	c.JSON(200, historyList)
}

// GetHistory @Summary GetHistory
// @Description Get History by CustomerId
// @Tags History
// @Produce application/json
// @Param CustomerId query string true "Customer id"
// @Success 200 {object} model.History
// @Failure 500 {string} string "{"Message": err.Error()}"
// @Router /history [get]
func (u *HistoryHandler) GetHistory(c *gin.Context) {
	customerId, _ := strconv.Atoi(c.Query("CustomerId"))
	historyData, err := u.ser.GetHistoryByID(customerId)
	if err != nil {
		if err.Error() == "error CRMS : There is no this customer" {
			c.JSON(200, gin.H{
				"Message": err.Error(),
			})
			return
		} else if err.Error() == "error CRMS : There is not any history" {
			c.JSON(200, gin.H{
				"Message": err.Error(),
			})
			return
		} else {
			c.JSON(500, gin.H{
				"Message": err.Error(),
			})
			return
		}
	}
	c.JSON(200, historyData)
}

// CreateHistory @Summary CreateHistory
// @Description Create a new History
// @Tags History
// @Produce application/json
// @Param History body model.History true "History Information"
// @Success 200 {object} model.History
// @Failure 500 {string} string "{"Message": err.Error()}"
// @Router /history [post]
func (u *HistoryHandler) CreateHistory(c *gin.Context) {
	json := model.HistoryRequest{}
	if err := c.BindJSON(&json); err != nil {
		c.JSON(400, gin.H{
			"Message": err.Error(),
		})
		return
	}
	createHistory := transformToHistory(json)
	createHistory, err := u.ser.CreateHistory(createHistory)
	if err != nil {
		if err.Error() == "error CRMS : There is no this customer" {
			c.JSON(200, gin.H{
				"Message": err.Error(),
			})
			return
		} else if err.Error() == "error CRMS : History Info is incomplete" {
			c.JSON(200, gin.H{
				"Message": err.Error(),
			})
			return
		} else {
			c.JSON(500, gin.H{
				"Message": err.Error(),
			})
			return
		}
	}
	c.JSON(200, createHistory)
}

// ModifyHistory @Summary ModifyHistory
// @Description Modify History
// @Tags History
// @Accept json
// @Produce application/json
// @Param History body model.History true "History Information"
// @Success 200 {object} model.History
// @Failure 500 {string} string "{"Message": err.Error()}"
// @Router /history [put]
func (u *HistoryHandler) ModifyHistory(c *gin.Context) {
	json := model.HistoryRequest{}
	if err := c.BindJSON(&json); err != nil {
		c.JSON(400, gin.H{
			"Message": err.Error(),
		})
		return
	}
	modifyHistory := transformToHistory(json)
	modifyHistory, err := u.ser.UpdateHistory(modifyHistory)
	if err != nil {
		if err.Error() == "error CRMS : There is no this history" {
			c.JSON(200, gin.H{
				"Message": err.Error(),
			})
			return
		} else if err.Error() == "error CRMS : Wrong customer" {
			c.JSON(200, gin.H{
				"Message": err.Error(),
			})
			return
		} else if err.Error() == "error CRMS : History Info is incomplete" {
			c.JSON(200, gin.H{
				"Message": err.Error(),
			})
			return
		} else {
			c.JSON(500, gin.H{
				"Message": err.Error(),
			})
			return
		}
	}
	c.JSON(200, modifyHistory)
}

// DeleteHistory @Summary DeleteHistory
// @Description Delete History by HistoryId
// @Tags History
// @Produce application/json
// @Param HistoryId query string true "History id"
// @Success 200 {object} string "Message": "Delete success"
// @Failure 500 {string} string "{"Message": err.Error()}"
// @Router /history [delete]
func (u *HistoryHandler) DeleteHistory(c *gin.Context) {
	historyId, _ := strconv.Atoi(c.Query("HistoryId"))
	err := u.ser.DeleteHistory(historyId)
	if err != nil {
		if err.Error() == "error CRMS : There is no this history" {
			c.JSON(200, gin.H{
				"Message": err.Error(),
			})
			return
		} else {
			c.JSON(500, gin.H{
				"Message": err.Error(),
			})
			return
		}
	}
	c.JSON(200, gin.H{
		"Message": "Delete success",
	})
}

func transformToHistory(requestData model.HistoryRequest) *model.History {
	date, _ := time.Parse("2006-01-02", requestData.Date)
	h := &model.History{
		CustomerId:     requestData.CustomerId,
		Date:           date,
		NumberOfPeople: requestData.NumberOfPeople,
		Price:          requestData.Price,
		Room:           requestData.Room,
		Note:           requestData.Note,
	}
	if requestData.HistoryId != 0 {
		h.HistoryId = requestData.HistoryId
	}
	return h
}
