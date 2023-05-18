package routes

import (
	"github.com/rogeriofontes/api-go-gin/api/controller"
	"github.com/rogeriofontes/api-go-gin/infrastructure"
)

//UserRoute -> Route for user module
type UserRoute struct {
	Handler    infrastructure.GinRouter
	Controller controller.UserController
}

//NewUserRoute -> initializes new instance of UserRoute
func NewUserRoute(
	controller controller.UserController,
	handler infrastructure.GinRouter,
) UserRoute {
	return UserRoute{
		Handler:    handler,
		Controller: controller,
	}
}

//Setup -> setups user routes
func (u UserRoute) Setup() {
	user := u.Handler.Gin.Group("/api/v1/auth")
	{
		user.POST("/register", u.Controller.CreateUser)
		user.POST("/login", u.Controller.LoginUser)
	}
}
