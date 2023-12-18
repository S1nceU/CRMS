package http

import (
	"github.com/S1nceU/CRMS/model"
	"github.com/S1nceU/CRMS/module/history"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"net/http"
	"time"
)

type HistoryHandler struct {
	ser history.Service
}

func NewHistoryHandler(e *gin.Engine, ser history.Service) {
	handler := &HistoryHandler{
		ser: ser,
	}
	api := e.Group("/api")
	{
		api.GET("/historyList", handler.ListHistories)
		api.GET("/history", handler.GetHistory)
		api.POST("/history", handler.CreateHistory)
		api.PUT("/history", handler.ModifyHistory)
		api.DELETE("/history", handler.DeleteHistory)
		api.POST("/historyForDuring", handler.GetHistoryForDuring)
	}

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
		c.JSON(http.StatusInternalServerError, gin.H{
			"Message": "Internal Error!",
		})
		return
	}
	c.JSON(http.StatusOK, historyList)
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
	customerId := uuid.MustParse(c.Query("CustomerId"))
	historyData, err := u.ser.GetHistoryByCustomerId(customerId)
	if err != nil {
		if err.Error() == "error CRMS : There is no this customer" {
			c.JSON(http.StatusOK, gin.H{
				"Message": err.Error(),
			})
			return
		} else if err.Error() == "error CRMS : There is not any history" {
			c.JSON(http.StatusOK, gin.H{
				"Message": err.Error(),
			})
			return
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{
				"Message": err.Error(),
			})
			return
		}
	}
	c.JSON(http.StatusOK, historyData)
}

// CreateHistory @Summary CreateHistory
// @Description Create a new History
// @Tags History
// @Produce application/json
// @Param History body model.HistoryRequest true "History Information"
// @Success 200 {object} model.History
// @Failure 500 {string} string "{"Message": err.Error()}"
// @Router /history [post]
func (u *HistoryHandler) CreateHistory(c *gin.Context) {
	json := model.HistoryRequest{}
	if err := c.BindJSON(&json); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"Message": err.Error(),
		})
		return
	}
	var err error
	createHistory, err := transformToHistory(json)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"Message": err.Error(),
		})
		return
	}
	createHistory, err = u.ser.CreateHistory(createHistory)
	if err != nil {
		if err.Error() == "error CRMS : There is no this customer" {
			c.JSON(http.StatusOK, gin.H{
				"Message": err.Error(),
			})
			return
		} else if err.Error() == "error CRMS : History Info is incomplete" {
			c.JSON(http.StatusOK, gin.H{
				"Message": err.Error(),
			})
			return
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{
				"Message": err.Error(),
			})
			return
		}
	}
	c.JSON(http.StatusCreated, createHistory)
}

// ModifyHistory @Summary ModifyHistory
// @Description Modify History
// @Tags History
// @Accept json
// @Produce application/json
// @Param History body model.HistoryRequest true "History Information"
// @Success 200 {object} model.History
// @Failure 500 {string} string "{"Message": err.Error()}"
// @Router /history [put]
func (u *HistoryHandler) ModifyHistory(c *gin.Context) {
	json := model.HistoryRequest{}
	if err := c.BindJSON(&json); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"Message": err.Error(),
		})
		return
	}
	var err error
	modifyHistory, err := transformToHistory(json)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"Message": err.Error(),
		})
		return
	}
	modifyHistory, err = u.ser.UpdateHistory(modifyHistory)
	if err != nil {
		if err.Error() == "error CRMS : There is no this history" {
			c.JSON(http.StatusOK, gin.H{
				"Message": err.Error(),
			})
			return
		} else if err.Error() == "error CRMS : Wrong customer" {
			c.JSON(http.StatusOK, gin.H{
				"Message": err.Error(),
			})
			return
		} else if err.Error() == "error CRMS : History Info is incomplete" {
			c.JSON(http.StatusOK, gin.H{
				"Message": err.Error(),
			})
			return
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{
				"Message": err.Error(),
			})
			return
		}
	}
	c.JSON(http.StatusOK, modifyHistory)
}

// DeleteHistory @Summary DeleteHistory
// @Description Delete History by HistoryId
// @Tags History
// @Produce application/json
// @Param HistoryId query uuid.UUID true "History id"
// @Success 200 {object} string "Message": "Delete success"
// @Failure 500 {string} string "{"Message": err.Error()}"
// @Router /history [delete]
func (u *HistoryHandler) DeleteHistory(c *gin.Context) {
	historyId := uuid.MustParse(c.Query("HistoryId"))
	err := u.ser.DeleteHistory(historyId)
	if err != nil {
		if err.Error() == "error CRMS : There is no this history" {
			c.JSON(http.StatusOK, gin.H{
				"Message": err.Error(),
			})
			return
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{
				"Message": err.Error(),
			})
			return
		}
	}
	c.JSON(http.StatusOK, gin.H{
		"Message": "Delete success",
	})
}

// GetHistoryForDuring @Summary GetHistoryForDuring
// @Description Get History For During
// @Tags History
// @Produce application/json
// @Param History body model.DuringRequest true "History Information"
// @Success 200 {object} model.History
// @Failure 500 {string} string "{"Message": err.Error()}"
// @Router /historyForDuring [post]
func (u *HistoryHandler) GetHistoryForDuring(c *gin.Context) {
	json := model.DuringRequest{}
	if err := c.BindJSON(&json); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"Message": err.Error(),
		})
		return
	}

	duringHistory, err := u.ser.GetHistoryForDuring(json.StartDate, json.EndDate)
	if err != nil {
		if err.Error() == "error CRMS : There is not any history between "+json.StartDate+" to "+json.EndDate {
			c.JSON(http.StatusOK, gin.H{
				"Message": err.Error(),
			})
			return
		} else if err.Error() == "error CRMS : Date is incomplete" {
			c.JSON(http.StatusOK, gin.H{
				"Message": err.Error(),
			})
			return
		} else if err.Error() == "error CRMS : Start date is after end date" {
			c.JSON(http.StatusOK, gin.H{
				"Message": err.Error(),
			})
			return
		} else if err.Error() == "error CRMS : End date is after today" {
			c.JSON(http.StatusOK, gin.H{
				"Message": err.Error(),
			})
			return
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{
				"Message": err.Error(),
			})
			return
		}
	}
	c.JSON(http.StatusOK, duringHistory)
}

func transformToHistory(requestData model.HistoryRequest) (*model.History, error) {
	date, err := time.ParseInLocation("2006-01-02", requestData.Date, time.Local)
	if err != nil {
		return nil, err
	}
	h := &model.History{
		CustomerId:     requestData.CustomerId,
		Date:           date,
		NumberOfPeople: requestData.NumberOfPeople,
		Price:          requestData.Price,
		Room:           requestData.Room,
		Note:           requestData.Note,
	}
	if requestData.HistoryId != uuid.Nil {
		h.HistoryId = requestData.HistoryId
	}
	return h, nil
}
