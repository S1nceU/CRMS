package http

import (
	"github.com/S1nceU/CRMS/model"
	"github.com/S1nceU/CRMS/module/customer"
	"github.com/S1nceU/CRMS/module/history"
	"github.com/gin-gonic/gin"
	"strconv"
	"time"
)

type CustomerHandler struct {
	customerSer customer.Service
	historySer  history.Service
}

func NewCustomerHandler(e *gin.Engine, customerSer customer.Service, historySer history.Service) {
	handler := &CustomerHandler{
		customerSer: customerSer,
		historySer:  historySer,
	}
	e.GET("/api/customerList", handler.ListCustomers)
	e.GET("/api/customer", handler.GetCustomerByID)
	e.POST("/api/customer", handler.CreateCustomer)
	e.PUT("/api/customer", handler.ModifyCustomer)
	e.DELETE("/api/customer", handler.DeleteCustomer)
}

// ListCustomers @Summary ListCustomers
// @Description Get all Customer
// @Accept json
// @Tags Customer
// @Produce application/json
// @Success 200 {object} model.Customer
// @Failure 500 {string} string "{"Message": "Internal Error!"}"
// @Router /customerList [get]
func (u *CustomerHandler) ListCustomers(c *gin.Context) {
	customerList, err := u.customerSer.ListCustomers()
	if err != nil {
		c.JSON(500, gin.H{
			"Message": "Internal Error!",
		})
		return
	}
	c.JSON(200, customerList)
}

// GetCustomerByID @Summary GetCustomerByID
// @Description Get Customer by ID
// @Tags Customer
// @Produce application/json
// @Param ID query string true "Customer ID" example(L123546789)
// @Success 200 {object} model.Customer
// @Failure 500 {string} string "{"Message": err.Error()}"
// @Router /customer [get]
func (u *CustomerHandler) GetCustomerByID(c *gin.Context) {
	iD := c.Query("ID")
	customerData, err := u.customerSer.GetCustomerByID(iD)
	if err != nil {
		if err.Error() == "error CRMS : There is no this customer" {
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
	c.JSON(200, customerData)
}

// CreateCustomer @Summary CreateCustomer
// @Description Create a new Customer
// @Tags Customer
// @Accept json
// @Produce application/json
// @Param Customer body model.Customer true "Customer Information"
// @Success 200 {object} model.Customer
// @Failure 500 {string} string "{"Message": err.Error()}"
// @Router /customer [post]
func (u *CustomerHandler) CreateCustomer(c *gin.Context) {
	json := model.CustomerRequest{}
	if err := c.BindJSON(&json); err != nil {
		c.JSON(400, gin.H{
			"Message": err.Error(),
		})
		return
	}
	createCustomer := transformToCustomer(json)
	createCustomer, err := u.customerSer.CreateCustomer(createCustomer)
	if err != nil {
		if err.Error() == "error CRMS : This customer is already existed" {
			c.JSON(200, gin.H{
				"Message": err.Error(),
			})
			return
		} else if err.Error() == "error CRMS : Customer Info is incomplete" {
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
	c.JSON(200, createCustomer)
}

// ModifyCustomer @Summary ModifyCustomer
// @Description Modify Customer
// @Tags Customer
// @Accept json
// @Produce application/json
// @Param Customer body model.Customer true "Customer Information"
// @Success 200 {object} model.Customer
// @Failure 500 {string} string "{"Message": err.Error()}"
// @Router /customer [put]
func (u *CustomerHandler) ModifyCustomer(c *gin.Context) {
	json := model.CustomerRequest{}
	if err := c.BindJSON(&json); err != nil {
		c.JSON(400, gin.H{
			"Message": err.Error(),
		})
		return
	}
	modifyCustomer := transformToCustomer(json)
	modifyCustomer, err := u.customerSer.UpdateCustomer(modifyCustomer)
	if err != nil {
		if err.Error() == "error CRMS : There is no this customer" {
			c.JSON(200, gin.H{
				"Message": err.Error(),
			})
			return
		} else if err.Error() == "error CRMS : Customer Info is incomplete" {
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
		"Customer info": modifyCustomer,
		"Message":       "Modify success",
	})
}

// DeleteCustomer @Summary DeleteCustomer
// @Description Delete Customer by CustomerId
// @Tags Customer
// @Produce application/json
// @Param CustomerId query string true "Customer id"
// @Success 200 {object} string "Message": "Delete success"
// @Failure 500 {string} string "{"Message": err.Error()}"
// @Router /customer [delete]
func (u *CustomerHandler) DeleteCustomer(c *gin.Context) {
	var err error
	customerId, _ := strconv.Atoi(c.Query("CustomerId"))
	err = u.historySer.DeleteHistoriesByCustomer(customerId)
	err = u.customerSer.DeleteCustomer(customerId)
	if err != nil {
		if err.Error() == "error CRMS : There is no this customer" {
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

func transformToCustomer(requestData model.CustomerRequest) *model.Customer {
	birthday, _ := time.Parse("2006-01-02", requestData.Birthday)
	c := &model.Customer{
		Name:        requestData.Name,
		Gender:      requestData.Gender,
		Birthday:    birthday,
		ID:          requestData.ID,
		Address:     requestData.Address,
		PhoneNumber: requestData.PhoneNumber,
		CarNumber:   requestData.CarNumber,
		Citizenship: requestData.Citizenship,
		Note:        requestData.Note,
	}
	if requestData.CustomerId != 0 {
		c.CustomerId = requestData.CustomerId
	}
	return c
}
