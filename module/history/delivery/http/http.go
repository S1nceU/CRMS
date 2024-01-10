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
		api.POST("/historyByHistoryId", handler.GetHistory)
		api.POST("/history", handler.CreateHistory)
		api.PUT("/history", handler.ModifyHistory)
		api.DELETE("/history", handler.DeleteHistory)
		api.POST("/historyForDuring", handler.GetHistoryForDuring)
		api.POST("/historyForDate", handler.GetHistoriesForDate)
		api.POST("/historyByCustomerId", handler.GetHistoryByCustomerId)
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
// @Description Get History by HistoryId
// @Tags History
// @Produce application/json
// @Param HistoryId body model.HistoryIdRequest true "History id" example: "f1b9d7c0-9f0f-4f1a-8f1a-4f1a9f0f4f1a"
// @Success 200 {object} model.History
// @Failure 500 {string} string "{"Message": err.Error()}"
// @Router /historyByHistoryId [post]
func (u *HistoryHandler) GetHistory(c *gin.Context) {
	request := model.HistoryIdRequest{}
	if err := c.BindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"Message": err.Error(),
		})
		return
	}
	historyData, err := u.ser.GetHistoryByHistoryId(request.HistoryId)
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
	c.JSON(http.StatusOK, historyData)
}

// CreateHistory @Summary CreateHistory
// @Description Create a new History
// @Tags History
// @Produce application/json
// @Param History body model.HistoryRequest true "History Information" example: {"CustomerId": "00000000-0000-0000-0000-000000000000", "Date": "2020-01-01", "NumberOfPeople": 1, "Price": 1000, "Room": "101", "Note": "test"}
// @Success 200 {object} model.History
// @Failure 500 {string} string "{"Message": err.Error()}"
// @Router /history [post]
func (u *HistoryHandler) CreateHistory(c *gin.Context) {
	request := model.HistoryRequest{}
	if err := c.BindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"Message": err.Error(),
		})
		return
	}
	var err error
	createHistory, err := transformToHistory(request)
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
// @Param History body model.HistoryRequest true "History Information" example: {"HistoryId": "00000000-0000-0000-0000-000000000000", "CustomerId": "00000000-0000-0000-0000-000000000000", "Date": "2020-01-01", "NumberOfPeople": 1, "Price": 10000, "Room": "101", "Note": "test"}
// @Success 200 {object} model.History
// @Failure 500 {string} string "{"Message": err.Error()}"
// @Router /history [put]
func (u *HistoryHandler) ModifyHistory(c *gin.Context) {
	request := model.HistoryRequest{}
	if err := c.BindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"Message": err.Error(),
		})
		return
	}
	var err error
	modifyHistory, err := transformToHistory(request)
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
// @Param HistoryId query string true "History id"
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
// @Param History body model.DuringRequest true "History Information" example: {"startDate": "2020-01-01", "endDate": "2020-01-02"}
// @Success 200 {object} model.History
// @Failure 500 {string} string "{"Message": err.Error()}"
// @Router /historyForDuring [post]
func (u *HistoryHandler) GetHistoryForDuring(c *gin.Context) {
	request := model.DuringRequest{}
	if err := c.BindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"Message": err.Error(),
		})
		return
	}

	duringHistory, err := u.ser.ListHistoriesForDuring(request.StartDate, request.EndDate)
	if err != nil {
		if err.Error() == "error CRMS : There is not any history between "+request.StartDate+" to "+request.EndDate {
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

// GetHistoriesForDate @Summary GetHistoriesForDate
// @Description Get Histories For Date
// @Tags History
// @Produce application/json
// @Param History body model.DateRequest true "History Information" example: {"Date": "2020-01-01"}
// @Success 200 {object} model.History
// @Failure 500 {string} string "Message": err.Error()"
// @Router /historyForDate [post]
func (u *HistoryHandler) GetHistoriesForDate(c *gin.Context) {
	request := model.DateRequest{}
	if err := c.BindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"Message": err.Error(),
		})
		return
	}
	historyList, err := u.ser.ListHistoriesForDate(request.Date)
	if err != nil {
		if err.Error() == "error CRMS : There is not any history" {
			c.JSON(http.StatusOK, gin.H{
				"Message": err.Error(),
			})
			return
		} else if err.Error() == "error CRMS : Date is incomplete" {
			c.JSON(http.StatusOK, gin.H{
				"Message": err.Error(),
			})
			return
		} else if err.Error() == "error CRMS : Date is after today" {
			c.JSON(http.StatusOK, gin.H{
				"Message": err.Error(),
			})
			return
		} else if err.Error() == "error CRMS : There was no customer in "+request.Date {
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
	c.JSON(http.StatusOK, historyList)
}

// GetHistoryByCustomerId @Summary GetHistoryByCustomerId
// @Description Get History By CustomerId
// @Tags History
// @Produce application/json
// @Param History body model.CustomerIdRequest true "History Information" example: {"CustomerId": "00000000-0000-0000-0000-000000000000"}
// @Success 200 {object} model.History
// @Failure 500 {string} string "Message": err.Error()"
// @Router /historyByCustomerId [post]
func (u *HistoryHandler) GetHistoryByCustomerId(c *gin.Context) {
	request := model.CustomerIdRequest{}
	if err := c.BindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"Message": err.Error(),
		})
		return
	}
	historyData, err := u.ser.ListHistoriesByCustomerId(request.CustomerId)
	if err != nil {
		if err.Error() == "error CRMS : There is no this customer" {
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
