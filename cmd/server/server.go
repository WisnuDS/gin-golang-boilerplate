package server

import (
	"net/http"

	"boilerplate.com/v1/configs"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

type Engine struct {
	Gin *gin.Engine
}

func InitEngine(env configs.Env) Engine {
	gin.SetMode(env.GinMode)
	engine := gin.Default()

	if env.Environment != "DEV" {
		engine.Use(cors.Default())
	} else {
		engine.Use(cors.New(cors.Config{
			AllowOrigins:     []string{"*"},
			AllowMethods:     []string{"PUT", "PATCH", "GET", "POST", "OPTIONS", "DELETE"},
			AllowHeaders:     []string{"*"},
			AllowCredentials: true,
		}))
	}

	engine.GET(env.PublicRoute+"/healt-check", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{"data": "API is running and accessible"})
	})

	/* Swagger init */
	// docs.SwaggerInfo.BasePath = constants.PUBLIC_ROUTE
	// docs.SwaggerInfo.Title = "Arvis Backoffice Service"
	// docs.SwaggerInfo.Description = "Arvis Backoffice Service API Documentation for Web."
	// docs.SwaggerInfo.Schemes = []string{env.Schema}
	// httpRouter.GET(constants.PUBLIC_ROUTE+"/docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	return Engine{
		Gin: engine,
	}
}
