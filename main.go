package main

import (
	"fmt"
	"log"
	"net/http"
	"svc-users-go/config"

	UserHandlerPackage "svc-users-go/module/v1/user/presenter"
	UserRepoPackage "svc-users-go/module/v1/user/repository"
	UserUseCasePackage "svc-users-go/module/v1/user/usecase"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func main() {
	fmt.Println("Server is running :)")

	mongoConnection, errorMongoConn := config.MongoConnection()

	if errorMongoConn != nil {
		log.Println("Error when connect mongo : ", errorMongoConn.Error())
	}

	echoRouter := echo.New()
	echoRouter.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{http.MethodGet, http.MethodPut, http.MethodPost, http.MethodDelete},
	}))

	// User modules
	userRepo := UserRepoPackage.NewUserRepository(mongoConnection)
	userUseCase := UserUseCasePackage.NewUserUseCase(userRepo)
	UserHandlerPackage.NewUserHandler(echoRouter, userUseCase)

	//Configuration of logger
	echoRouter.Use(middleware.Logger())
	echoRouter.Logger.Fatal(echoRouter.Start(":8081"))
}
