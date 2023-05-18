package infrastructure

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/rogeriofontes/api-go-gin/docs"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

//GinRouter -> Gin Router
type GinRouter struct {
	Gin *gin.Engine
}

func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {

		c.Header("Access-Control-Allow-Origin", "*")
		c.Header("Access-Control-Allow-Credentials", "true")
		c.Header("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Header("Access-Control-Allow-Methods", "POST,HEAD,PATCH, OPTIONS, GET, PUT")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	}
}

// NewGinRouter all the routes are defined here

// NewGinRouter godoc
// @Summary Show the status of server.
// @Description get the status of server.
// @Tags root
// @Accept */*
// @Produce json
// @Success 200 {object} map[string]interface{}
// @Router /healthcheck [get]
func NewGinRouter() GinRouter {

	httpRouter := gin.Default()
	httpRouter.Use(CORSMiddleware())

	docs.SwaggerInfo.BasePath = "/api/v1"

	httpRouter.GET("/api/v1/healthcheck", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"data": "Polls Up and Running..."})
	})

	//url := ginSwagger.URL("http://localhost:8000/docs/swagger.json")

	//url := ginSwagger.URL("http://localhost:8000/docs/swagger.json") // The url pointing to API definition
	//httpRouter.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler, url))
	httpRouter.GET("/docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	return GinRouter{
		Gin: httpRouter,
	}

}
