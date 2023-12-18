package main

import (
	"fmt"
	"github.com/S1nceU/CRMS/config"
	_ "github.com/S1nceU/CRMS/docs"
	"github.com/S1nceU/CRMS/model"
	_customerHandlerHttpDelivery "github.com/S1nceU/CRMS/module/customer/delivery/http"
	_historyHandlerHttpDelivery "github.com/S1nceU/CRMS/module/history/delivery/http"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"strconv"

	_customerRepo "github.com/S1nceU/CRMS/module/customer/repository"
	_customerSer "github.com/S1nceU/CRMS/module/customer/service"

	_historyRepo "github.com/S1nceU/CRMS/module/history/repository"
	_historySer "github.com/S1nceU/CRMS/module/history/service"

	"github.com/S1nceU/CRMS/route"
)

var swagHandler gin.HandlerFunc

// @title CRMS_Swagger
// @version 1.0
// @description CRMS_Swagger information
// @termsOfService http://www.google.com

// @contact.name Jason Yang
// @contact.url http://www.google.com
// @contact.email jjkk900925@gmail.com

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host 127.0.0.1:8080
// @BasePath /api

func init() {
	swagHandler = ginSwagger.WrapHandler(swaggerFiles.Handler)
	config.Init()
}

func main() {
	var (
		db    *gorm.DB
		dbErr error
		dsn   string
	)
	dsn = fmt.Sprintf("%s:%s@%s(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		config.Val.DatabaseConfig.Username,
		config.Val.DatabaseConfig.Password,
		config.Val.DatabaseConfig.Network,
		config.Val.DatabaseConfig.Server,
		config.Val.DatabaseConfig.Port,
		config.Val.DatabaseConfig.Database,
	)
	if db, dbErr = gorm.Open(mysql.Open(dsn), &gorm.Config{}); dbErr != nil {
		panic("使用 gorm 連線 DB 發生錯誤，原因為 " + dbErr.Error())
	} else {
		fmt.Println("Connect to DB successfully")
		var err error
		if err = db.AutoMigrate(&model.Customer{}); err != nil {
			return
		}
		if err = db.AutoMigrate(&model.History{}); err != nil {
			return
		}
	}

	gin.SetMode(config.Val.Mode)
	server := gin.Default()

	customerRepo := _customerRepo.NewCustomerRepository(db)
	historyRepo := _historyRepo.NewHistoryRepository(db)

	customerSer := _customerSer.NewCustomerService(customerRepo)
	historySer := _historySer.NewHistoryService(historyRepo)

	_customerHandlerHttpDelivery.NewCustomerHandler(server, customerSer, historySer)
	_historyHandlerHttpDelivery.NewHistoryHandler(server, historySer)

	route.NewRoute(server)

	if swagHandler != nil {
		server.GET("swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	}
	err := server.Run(":" + strconv.Itoa(config.Val.Port))
	if err != nil {
		return
	}
}
