package main

import (
	"context"
	"errors"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"strconv"
	"time"

	"github.com/S1nceU/CRMS/apps/api/config"
	_ "github.com/S1nceU/CRMS/apps/api/docs"
	"github.com/S1nceU/CRMS/apps/api/model"

	_citizenshipHandlerHttpDelivery "github.com/S1nceU/CRMS/apps/api/module/citizenship/delivery/http"
	_citizenshipRepo "github.com/S1nceU/CRMS/apps/api/module/citizenship/repository"
	_citizenshipSer "github.com/S1nceU/CRMS/apps/api/module/citizenship/service"
	_customerHandlerHttpDelivery "github.com/S1nceU/CRMS/apps/api/module/customer/delivery/http"
	_customerRepo "github.com/S1nceU/CRMS/apps/api/module/customer/repository"
	_customerSer "github.com/S1nceU/CRMS/apps/api/module/customer/service"
	_historyHandlerHttpDelivery "github.com/S1nceU/CRMS/apps/api/module/history/delivery/http"
	_historyRepo "github.com/S1nceU/CRMS/apps/api/module/history/repository"
	_historySer "github.com/S1nceU/CRMS/apps/api/module/history/service"
	_userHandlerHttpDelivery "github.com/S1nceU/CRMS/apps/api/module/user/delivery/http"
	_userRepo "github.com/S1nceU/CRMS/apps/api/module/user/repository"
	_userSer "github.com/S1nceU/CRMS/apps/api/module/user/service"
	"github.com/S1nceU/CRMS/apps/api/route"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var swagHandler gin.HandlerFunc
var db *gorm.DB

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
	//swagHandler = ginSwagger.WrapHandler(swaggerFiles.Handler)
	config.Init()

	var (
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
		log.Fatal("There was an error connecting to the DB using gorm, due to " + dbErr.Error())
	} else {
		log.Println("Connect to DB successfully")
		var err error
		if !db.Migrator().HasTable("citizenships") {
			if err = db.AutoMigrate(&model.Citizenship{}); err != nil {
				return
			}
			config.ImportCitizenshipData(db)
			log.Println("Init citizenship data successfully")
		}
		if err = db.AutoMigrate(&model.Customer{}); err != nil {
			return
		}
		if err = db.AutoMigrate(&model.History{}); err != nil {
			return
		}
		if err = db.AutoMigrate(&model.User{}); err != nil {
			return
		}
	}
}

func main() {

	gin.SetMode(config.Val.Mode)
	router := gin.Default()
	router.Use(route.Cors()) // CORS middleware

	customerRepo := _customerRepo.NewCustomerRepository(db)
	historyRepo := _historyRepo.NewHistoryRepository(db)
	userRepo := _userRepo.NewUserRepository(db)
	citizenshipRepo := _citizenshipRepo.NewCitizenshipRepository(db)

	customerSer := _customerSer.NewCustomerService(customerRepo)
	historySer := _historySer.NewHistoryService(historyRepo)
	userSer := _userSer.NewUserService(userRepo)
	citizenshipSer := _citizenshipSer.NewCitizenshipService(citizenshipRepo)

	_customerHandlerHttpDelivery.NewCustomerHandler(router, customerSer)
	_historyHandlerHttpDelivery.NewHistoryHandler(router, historySer)
	_citizenshipHandlerHttpDelivery.NewCitizenshipHandler(router, citizenshipSer)
	_userHandlerHttpDelivery.NewUserHandler(router, userSer)

	route.NewRoute(router)

	if swagHandler != nil {
		router.GET("swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	}

	server := &http.Server{
		Addr:    ":" + strconv.Itoa(config.Val.Port),
		Handler: router,
	}

	log.Println("Server is running")
	go func() {
		if err := server.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			log.Fatalf("listen: %s\n", err)
		}
	}()

	quit := make(chan os.Signal)
	signal.Notify(quit, os.Interrupt)
	<-quit
	log.Println("Shutdown Server ...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := server.Shutdown(ctx); err != nil {
		log.Fatal("Server Shutdown:", err)
	}

	log.Println("Closing DB connection")
	sqlDB, _ := db.DB()
	if err := sqlDB.Close(); err != nil {
		log.Fatalf("Error closing DB connection: %v", err)
	}

	log.Println("Server exiting")
}
