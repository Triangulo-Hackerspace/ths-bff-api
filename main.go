package main

import (
	"github.com/rogeriofontes/api-go-gin/api/controller"
	"github.com/rogeriofontes/api-go-gin/api/repository"
	"github.com/rogeriofontes/api-go-gin/api/routes"
	"github.com/rogeriofontes/api-go-gin/api/service"
	"github.com/rogeriofontes/api-go-gin/database"
	"github.com/rogeriofontes/api-go-gin/infrastructure"
	"github.com/rogeriofontes/api-go-gin/models"
)

func init() {
	infrastructure.LoadEnv()
}

// @title Gin Swagger Example API
// @version 1.0
// @description This is a sample server server.
// @termsOfService http://swagger.io/terms/

// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host localhost:8000
// @BasePath /
// @schemes http

// @securityDefinitions.apikey Bearer
// @in header
// @name Authorization
// @description "Type 'Bearer TOKEN' to correctly set the API Key"
func main() {
	router := infrastructure.NewGinRouter()
	db := database.NewDatabase()

	postRepository := repository.NewPostRepository(db)
	postService := service.NewPostService(postRepository)
	postController := controller.NewPostController(postService)
	postRoute := routes.NewPostRoute(postController, router)
	postRoute.Setup()

	userRepository := repository.NewUserRepository(db)
	userService := service.NewUserService(userRepository)
	userController := controller.NewUserController(userService)
	userRoute := routes.NewUserRoute(userController, router)
	userRoute.Setup()

	commentRepository := repository.NewCommentRepository(db)
	commentService := service.NewCommentService(commentRepository)
	commentController := controller.NewCommentController(commentService)
	commentRoute := routes.NewCommentRoute(commentController, router)
	commentRoute.Setup()

	projectRepository := repository.NewProjectRepository(db)
	projectService := service.NewProjectService(projectRepository)
	projectController := controller.NewProjectController(projectService)
	projectRoute := routes.NewProjectRoute(projectController, router)
	projectRoute.Setup()

	eventRepository := repository.NewEventRepository(db)
	eventService := service.NewEventService(eventRepository)
	eventController := controller.NewEventController(eventService)
	eventRoute := routes.NewEventRoute(eventController, router)
	eventRoute.Setup()

	meetRepository := repository.NewMeetRepository(db)
	meetService := service.NewMeetService(meetRepository)
	meetController := controller.NewMeetController(meetService)
	meetRoute := routes.NewMeetRoute(meetController, router)
	meetRoute.Setup()

	profileRepository := repository.NewProfileRepository(db)
	profileService := service.NewProfileService(profileRepository)
	profileController := controller.NewProfileController(profileService)
	profileRoute := routes.NewProfileRoute(profileController, router)
	profileRoute.Setup()

	db.DB.AutoMigrate(&models.Post{}, &models.User{}, &models.Comment{}, &models.Project{}, &models.Event{}, &models.Meet{}, &models.Profile{})
	router.Gin.Run(":8000")
}
