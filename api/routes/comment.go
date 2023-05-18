package routes

import (
	"github.com/rogeriofontes/api-go-gin/api/controller"
	"github.com/rogeriofontes/api-go-gin/infrastructure"
	"github.com/rogeriofontes/api-go-gin/middlewares"
)

//CommentRoute -> Route for question module
type CommentRoute struct {
	Controller controller.CommentController
	Handler    infrastructure.GinRouter
}

//NewCommentRoute -> initializes new choice rouets
func NewCommentRoute(
	controller controller.CommentController,
	handler infrastructure.GinRouter,

) CommentRoute {
	return CommentRoute{
		Controller: controller,
		Handler:    handler,
	}
}

//Setup -> setups new choice Routes
func (p CommentRoute) Setup() {
	comment := p.Handler.Gin.Group("/api/v1/comments")
	comment.Use(middlewares.JwtAuthMiddleware())
	{
		comment.GET("/", p.Controller.GetComments)
		comment.POST("/", p.Controller.AddComment)
		comment.GET("/:id", p.Controller.GetComment)
		comment.DELETE("/:id", p.Controller.DeleteComment)
		comment.PUT("/:id", p.Controller.UpdateComment)
	}
}
