package http

import (
	"github.com/S1nceU/CRMS/domain"
	"github.com/S1nceU/CRMS/model"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"net/http"
	"time"
)

type CustomerHandler struct {
	customerSer domain.CustomerService
	historySer  domain.HistoryService // if frontend can block delete customer when this customer have history, we can remove this line
}

func NewCustomerHandler(e *gin.Engine, customerSer domain.CustomerService, historySer domain.HistoryService) {
	handler := &CustomerHandler{
		customerSer: customerSer,
		historySer:  historySer,
	}
	api := e.Group("/api")
	{
		api.GET("/customerList", handler.ListCustomers)
		api.GET("/customer", handler.GetCustomerByID)
		api.POST("/customer", handler.CreateCustomer)
		api.PUT("/customer", handler.ModifyCustomer)
		api.DELETE("/customer", handler.DeleteCustomer)
		api.POST("/customerName", handler.GetCustomerByCustomerName)
		api.GET("/customerCitizenship", handler.ListCustomersByCitizenship)
		api.GET("/customerPhone", handler.GetCustomerByCustomerPhone)
		api.POST("/customerID", handler.GetCustomerByCustomerID)
	}
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
		c.JSON(http.StatusInternalServerError, gin.H{
			"Message": "Internal Error!",
		})
		return
	}
	c.JSON(http.StatusOK, customerList)
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
	id := c.Query("ID")
	customerData, err := u.customerSer.GetCustomerByID(id)
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
	c.JSON(http.StatusOK, customerData)
}

// CreateCustomer @Summary CreateCustomer
// @Description Create a new Customer
// @Tags Customer
// @Accept json
// @Produce application/json
// @Param Customer body model.CustomerRequest true "Customer Information"
// @Success 200 {object} model.Customer
// @Failure 500 {string} string "{"Message": err.Error()}"
// @Router /customer [post]
func (u *CustomerHandler) CreateCustomer(c *gin.Context) {
	request := model.CustomerRequest{}
	if err := c.BindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"Message": err.Error(),
		})
		return
	}
	createCustomer, err := transformToCustomer(request)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"Message": err.Error(),
		})
		return
	}
	createCustomer, err = u.customerSer.CreateCustomer(createCustomer)
	if err != nil {
		if err.Error() == "error CRMS : This customer is already existed" {
			c.JSON(http.StatusOK, gin.H{
				"Message": err.Error(),
			})
			return
		} else if err.Error() == "error CRMS : Customer Info is incomplete" {
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
	c.JSON(http.StatusOK, createCustomer)
}

// ModifyCustomer @Summary ModifyCustomer
// @Description Modify Customer
// @Tags Customer
// @Accept json
// @Produce application/json
// @Param Customer body model.CustomerRequest true "Customer Information"
// @Success 200 {object} model.Customer
// @Failure 500 {string} string "{"Message": err.Error()}"
// @Router /customer [put]
func (u *CustomerHandler) ModifyCustomer(c *gin.Context) {
	request := model.CustomerRequest{}
	if err := c.BindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"Message": err.Error(),
		})
		return
	}
	modifyCustomer, err := transformToCustomer(request)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"Message": err.Error(),
		})
		return
	}
	modifyCustomer, err = u.customerSer.UpdateCustomer(modifyCustomer)
	if err != nil {
		if err.Error() == "error CRMS : There is no this customer" {
			c.JSON(http.StatusOK, gin.H{
				"Message": err.Error(),
			})
			return
		} else if err.Error() == "error CRMS : Customer Info is incomplete" {
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
		"Customer info": modifyCustomer,
		"Message":       "Modify success",
	})
}

// DeleteCustomer @Summary DeleteCustomer
// @Description Delete Customer by CustomerId
// @Tags Customer
// @Produce application/json
// @Param CustomerId query uuid.UUID true "Customer id"
// @Success 200 {object} string "Message": "Delete success"
// @Failure 500 {string} string "{"Message": err.Error()}"
// @Router /customer [delete]
func (u *CustomerHandler) DeleteCustomer(c *gin.Context) {
	var err error
	customerId := uuid.MustParse(c.Query("CustomerId"))
	err = u.historySer.DeleteHistoriesByCustomer(customerId)
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
	err = u.customerSer.DeleteCustomer(customerId)
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
	c.JSON(http.StatusOK, gin.H{
		"Message": "Delete success",
	})
}

// GetCustomerByCustomerName @Summary GetCustomerByCustomerName
// @Description Get Customer by CustomerName
// @Tags Customer
// @Produce application/json
// @Param CustomerName query string true "Customer name"
// @Success 200 {object} string "Message": "Delete success"
// @Failure 500 {string} string "{"Message": err.Error()}"
// @Router /customerName [post]
func (u *CustomerHandler) GetCustomerByCustomerName(c *gin.Context) {
	customerName := c.Query("CustomerName")
	customerData, err := u.customerSer.ListCustomersByCustomerName(customerName)
	if err != nil {
		if err.Error() == "error CRMS : Customer Info is incomplete" {
			c.JSON(http.StatusOK, gin.H{
				"Message": err.Error(),
			})
			return
		} else if err.Error() == "error CRMS : There is no this customer" {
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
	c.JSON(http.StatusOK, customerData)
}

// ListCustomersByCitizenship @Summary ListCustomersByCitizenship
// @Description Get all Customers by citizenship
// @Tags Customer
// @Produce application/json
// @Param Citizenship query string true "Citizenship" example(Taiwan)
// @Success 200 {object} []model.Customer
// @Failure 500 {string} string "{"Message": err.Error()}"
// @Router /customerCitizenship [get]
func (u *CustomerHandler) ListCustomersByCitizenship(c *gin.Context) {
	request := model.CustomerCitizenshipRequest{}
	if err := c.BindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"Message": err.Error(),
		})
		return
	}
	customerData, err := u.customerSer.ListCustomersByCitizenship(request.Citizenship)
	if err != nil {
		if err.Error() == "error CRMS : Customer Info is incomplete" {
			c.JSON(http.StatusOK, gin.H{
				"Message": err.Error(),
			})
			return
		} else if err.Error() == "error CRMS : There is no this customer" {
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
	c.JSON(http.StatusOK, customerData)
}

// GetCustomerByCustomerPhone @Summary GetCustomerByCustomerPhone
// @Description Get Customer by CustomerPhone
// @Tags Customer
// @Produce application/json
// @Param CustomerPhone query string true "Customer phone" example(0912345678)
// @Success 200 {object} []model.Customer
// @Failure 500 {string} string "{"Message": err.Error()}"
// @Router /customerPhone [get]
func (u *CustomerHandler) GetCustomerByCustomerPhone(c *gin.Context) {
	customerPhone := c.Query("CustomerPhone")
	customerData, err := u.customerSer.ListCustomersByCustomerPhone(customerPhone)
	if err != nil {
		if err.Error() == "error CRMS : Customer Info is incomplete" {
			c.JSON(http.StatusOK, gin.H{
				"Message": err.Error(),
			})
			return
		} else if err.Error() == "error CRMS : There is no this customer" {
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
	c.JSON(http.StatusOK, customerData)
}

// GetCustomerByCustomerID @Summary GetCustomerByCustomerID
// @Description Get Customer by CustomerID
// @Tags Customer
// @Produce application/json
// @Param CustomerId body model.CustomerIdRequest true "Customer id"
// @Success 200 {object} []model.Customer
// @Failure 500 {string} string "{"Message": err.Error()}"
// @Router /customerID [post]
func (u *CustomerHandler) GetCustomerByCustomerID(c *gin.Context) {
	request := model.CustomerIdRequest{}
	if err := c.BindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"Message": err.Error(),
		})
		return
	}

	customerData, err := u.customerSer.GetCustomerByCustomerId(request.CustomerId)
	if err != nil {
		if err.Error() == "error CRMS : Invalid request" {
			c.JSON(http.StatusOK, gin.H{
				"Message": err.Error(),
			})
			return
		} else if err.Error() == "error CRMS : There is no this customer" {
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
	c.JSON(http.StatusOK, customerData)
}

func transformToCustomer(requestData model.CustomerRequest) (*model.Customer, error) {
	birthday, err := time.ParseInLocation("2006-01-02", requestData.Birthday, time.Local)
	if err != nil {
		return nil, err
	}
	c := &model.Customer{
		Name:          requestData.Name,
		Gender:        requestData.Gender,
		Birthday:      birthday,
		ID:            requestData.ID,
		Address:       requestData.Address,
		PhoneNumber:   requestData.PhoneNumber,
		CarNumber:     requestData.CarNumber,
		CitizenshipId: requestData.Citizenship,
		Note:          requestData.Note,
		CustomerId:    requestData.CustomerId,
	}
	return c, nil
}
