package routes

import (
	"github.com/rogeriofontes/api-go-gin/api/controller"
	"github.com/rogeriofontes/api-go-gin/infrastructure"
	"github.com/rogeriofontes/api-go-gin/middlewares"
)

//ProjectRoute -> Route for question module
type ProjectRoute struct {
	Controller controller.ProjectController
	Handler    infrastructure.GinRouter
}

//NewProjectRoute -> initializes new choice rouets
func NewProjectRoute(
	controller controller.ProjectController,
	handler infrastructure.GinRouter,

) ProjectRoute {
	return ProjectRoute{
		Controller: controller,
		Handler:    handler,
	}
}

//Setup -> setups new choice Routes
func (p ProjectRoute) Setup() {
	project := p.Handler.Gin.Group("/api/v1/projects")
	project.Use(middlewares.JwtAuthMiddleware())
	{
		project.GET("/", p.Controller.GetProjects)
		project.POST("/", p.Controller.AddProject)
		project.GET("/:id", p.Controller.GetProject)
		project.DELETE("/:id", p.Controller.DeleteProject)
		project.PUT("/:id", p.Controller.UpdateProject)
	}
}
