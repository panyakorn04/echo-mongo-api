package routes

import (
	"github.com/labstack/echo/v4"
	"github.com/panyakorn4/echo-mongo-api/pkg/controllers"
)

func UserRoute(e *echo.Echo) {
	//All routes related to users comes here

	e.POST("/user", controllers.CreateUser)
	e.GET("/user/:userId", controllers.GetAUser)
	e.PATCH("/user/:userId", controllers.EditAUser)
	e.DELETE("/user/:userId", controllers.DeleteAUser)
	e.GET("/users", controllers.GetAllUsers)
}
