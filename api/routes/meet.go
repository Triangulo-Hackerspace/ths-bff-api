package routes

import (
	"github.com/rogeriofontes/api-go-gin/api/controller"
	"github.com/rogeriofontes/api-go-gin/infrastructure"
	"github.com/rogeriofontes/api-go-gin/middlewares"
)

//MeetRoute -> Route for question module
type MeetRoute struct {
	Controller controller.MeetController
	Handler    infrastructure.GinRouter
}

//NewMeetRoute -> initializes new choice rouets
func NewMeetRoute(
	controller controller.MeetController,
	handler infrastructure.GinRouter,

) MeetRoute {
	return MeetRoute{
		Controller: controller,
		Handler:    handler,
	}
}

//Setup -> setups new choice Routes
func (p MeetRoute) Setup() {
	meet := p.Handler.Gin.Group("/api/v1/meets")
	meet.Use(middlewares.JwtAuthMiddleware())
	{
		meet.GET("/", p.Controller.GetMeets)
		meet.POST("/", p.Controller.AddMeet)
		meet.GET("/:id", p.Controller.GetMeet)
		meet.DELETE("/:id", p.Controller.DeleteMeet)
		meet.PUT("/:id", p.Controller.UpdateMeet)
	}
}
