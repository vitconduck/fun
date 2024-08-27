package http

import (
	"log/slog"
	"strings"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	sloggin "github.com/samber/slog-gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/vitconduck/fun/pkg/configs"
)

type Router struct {
	*gin.Engine
}

func NewRouter(
	cfg *configs.HTTP,
	userHandler UserHandler,
) (*Router, error) {
	// Disable debug mode in production
	if cfg.Env == "production" {
		gin.SetMode(gin.ReleaseMode)
	}

	// CORS
	ginConfig := cors.DefaultConfig()
	allowedOrigins := cfg.AllowedOrigins
	originsList := strings.Split(allowedOrigins, ",")
	ginConfig.AllowOrigins = originsList

	rt := gin.New()
	rt.Use(sloggin.New(slog.Default()), gin.Recovery(), cors.New(ginConfig))

	rt.GET("/docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	v1 := rt.Group("/v1")
	{
		user := v1.Group("/users")
		{
			user.GET("/:id", userHandler.GetUser)
		}
	}

	return &Router{rt}, nil

}

// Serve starts the HTTP server
func (r *Router) Serve(listenAddr string) error {
	return r.Run(listenAddr)
}
