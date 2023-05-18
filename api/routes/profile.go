package routes

import (
	"github.com/rogeriofontes/api-go-gin/api/controller"
	"github.com/rogeriofontes/api-go-gin/infrastructure"
	"github.com/rogeriofontes/api-go-gin/middlewares"
)

//ProfileRoute -> Route for question module
type ProfileRoute struct {
	Controller controller.ProfileController
	Handler    infrastructure.GinRouter
}

//NewProfileRoute -> initializes new choice rouets
func NewProfileRoute(
	controller controller.ProfileController,
	handler infrastructure.GinRouter,

) ProfileRoute {
	return ProfileRoute{
		Controller: controller,
		Handler:    handler,
	}
}

//Setup -> setups new choice Routes
func (p ProfileRoute) Setup() {
	Profile := p.Handler.Gin.Group("/api/v1/profiles")
	Profile.Use(middlewares.JwtAuthMiddleware())
	{
		Profile.GET("/", p.Controller.GetProfiles)
		Profile.POST("/", p.Controller.AddProfile)
		Profile.GET("/:id", p.Controller.GetProfile)
		Profile.GET("/user/:id", p.Controller.GetProfileByUserId)
		Profile.DELETE("/:id", p.Controller.DeleteProfile)
		Profile.PUT("/:id", p.Controller.UpdateProfile)
	}
}
