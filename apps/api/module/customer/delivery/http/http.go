package http

import (
	"github.com/S1nceU/CRMS/domain"
	"github.com/S1nceU/CRMS/model"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

type CustomerHandler struct {
	customerSer domain.CustomerService
}

func NewCustomerHandler(e *gin.Engine, customerSer domain.CustomerService) {
	handler := &CustomerHandler{
		customerSer: customerSer,
	}
	api := e.Group("/api")
	{
		api.POST("/customerList", handler.ListCustomers)
		api.POST("/customerNationalId", handler.GetCustomerByNationalId)
		api.POST("/customerCre", handler.CreateCustomer)
		api.POST("/customerMod", handler.ModifyCustomer)
		api.POST("/customerDel", handler.DeleteCustomer)
		api.POST("/customerName", handler.GetCustomerByCustomerName)
		api.POST("/customerCitizenship", handler.ListCustomersByCitizenship)
		api.POST("/customerPhone", handler.GetCustomerByCustomerPhone)
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
// @Router /customerList [post]
func (u *CustomerHandler) ListCustomers(c *gin.Context) {
	customerList, err := u.customerSer.ListCustomers()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"Message": "Internal Error!",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"Message":   "List all customers",
		"customers": customerList,
	})
}

// GetCustomerByNationalId @Summary GetCustomerByNationalId
// @Description Get Customer by ID
// @Tags Customer
// @Produce application/json
// @Param ID body model.CustomerNationalIdRequest true "National ID"
// @Success 200 {object} model.Customer
// @Failure 500 {string} string "{"Message": err.Error()}"
// @Router /customerNationalId [post]
func (u *CustomerHandler) GetCustomerByNationalId(c *gin.Context) {
	request := model.CustomerNationalIdRequest{}
	if err := c.BindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"Message": err.Error(),
		})
		return
	}
	customerData, err := u.customerSer.GetCustomerByNationalId(request.NationalId)
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
// @Router /customerCre [post]
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
// @Router /customerMod [post]
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
// @Param CustomerId body model.CustomerIdRequest true "Customer ID"
// @Success 200 {object} string "Message": "Delete success"
// @Failure 500 {string} string "{"Message": err.Error()}"
// @Router /customerDel [post]
func (u *CustomerHandler) DeleteCustomer(c *gin.Context) {
	request := model.CustomerIdRequest{}
	if err := c.BindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"Message": err.Error(),
		})
		return
	}
	err := u.customerSer.DeleteCustomer(request.CustomerId)
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
// @Param Name body model.CustomerNameRequest true "Customer Name"
// @Success 200 {object} string "Message": "Delete success"
// @Failure 500 {string} string "{"Message": err.Error()}"
// @Router /customerName [post]
func (u *CustomerHandler) GetCustomerByCustomerName(c *gin.Context) {
	request := model.CustomerNameRequest{}
	if err := c.BindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"Message": err.Error(),
		})
		return
	}
	customerData, err := u.customerSer.ListCustomersByName(request.Name)
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
// @Param Citizenship body model.CustomerCitizenshipRequest true "Citizenship"
// @Success 200 {object} []model.Customer
// @Failure 500 {string} string "{"Message": err.Error()}"
// @Router /customerCitizenship [post]
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
// @Param CustomerPhone body model.CustomerPhoneRequest true "Customer Phone"
// @Success 200 {object} []model.Customer
// @Failure 500 {string} string "{"Message": err.Error()}"
// @Router /customerPhone [post]
func (u *CustomerHandler) GetCustomerByCustomerPhone(c *gin.Context) {
	request := model.CustomerPhoneRequest{}
	if err := c.BindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"Message": err.Error(),
		})
		return
	}
	customerData, err := u.customerSer.ListCustomersByPhone(request.PhoneNumber)
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
		NationalId:    requestData.NationalId,
		Address:       requestData.Address,
		PhoneNumber:   requestData.PhoneNumber,
		CarNumber:     requestData.CarNumber,
		CitizenshipId: requestData.Citizenship,
		Note:          requestData.Note,
		Id:            requestData.CustomerId,
	}
	return c, nil
}
