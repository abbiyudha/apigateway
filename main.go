package main

import (
	"ApiGateway/handler"
	"ApiGateway/middleware"
	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()
	//Endpoint Admin
	e.GET("/gateway/admin", handler.AdminProfile(), middleware.JWTMiddleware())
	e.POST("/gateway/create/admin", handler.CreateAdmin())
	e.POST("/gateway/admin/login", handler.Login())
	e.PUT("/gateway/update", handler.UpdateAdmin(), middleware.JWTMiddleware())
	e.DELETE("/gateway/delete", handler.DeleteAdmin(), middleware.JWTMiddleware())
	e.POST("/gateway/create/user", handler.CreateUser(), middleware.JWTMiddleware())

	//Endpoint User
	e.POST("/gateway/user/login", handler.LoginUser())
	e.GET("/gateway/user", handler.UserProfile(), middleware.JWTMiddleware())

	e.Start(":8000")
}
