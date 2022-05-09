package main

import (
	"ApiGateway/handler"
	"ApiGateway/middleware"
	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()
	//Endpoint Admin
	e.GET("/admin", handler.AdminProfile(), middleware.JWTMiddleware())
	e.POST("/create/admin", handler.CreateAdmin())
	e.POST("/admin/login", handler.Login())
	e.PUT("/update", handler.UpdateAdmin(), middleware.JWTMiddleware())
	e.DELETE("/delete", handler.DeleteAdmin(), middleware.JWTMiddleware())
	e.POST("/create/user", handler.CreateUser(), middleware.JWTMiddleware())

	//Endpoint User
	e.POST("/user/login", handler.LoginUser())
	e.GET("/user", handler.UserProfile(), middleware.JWTMiddleware())

	e.Start("localhost:8000")
}
