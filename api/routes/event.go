package routes

import (
	"github.com/rogeriofontes/api-go-gin/api/controller"
	"github.com/rogeriofontes/api-go-gin/infrastructure"
	"github.com/rogeriofontes/api-go-gin/middlewares"
)

//EventRoute -> Route for question module
type EventRoute struct {
	Controller controller.EventController
	Handler    infrastructure.GinRouter
}

//NewEventRoute -> initializes new choice rouets
func NewEventRoute(
	controller controller.EventController,
	handler infrastructure.GinRouter,

) EventRoute {
	return EventRoute{
		Controller: controller,
		Handler:    handler,
	}
}

//Setup -> setups new choice Routes
func (p EventRoute) Setup() {
	event := p.Handler.Gin.Group("/api/v1/events")
	event.Use(middlewares.JwtAuthMiddleware())
	{
		event.GET("/", p.Controller.GetEvents)
		event.POST("/", p.Controller.AddEvent)
		event.GET("/:id", p.Controller.GetEvent)
		event.DELETE("/:id", p.Controller.DeleteEvent)
		event.PUT("/:id", p.Controller.UpdateEvent)
	}
}
