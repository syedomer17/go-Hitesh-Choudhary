package httpserver

import (
	"go-auth/internal/app"
	"go-auth/internal/user"

	"github.com/gin-gonic/gin"
)

func NewRouter(a *app.App) *gin.Engine {

	router := gin.New()

	router.Use(gin.Logger())
	router.Use(gin.Recovery())

	router.GET("/health", health)

	userRepo := user.NewRepo(a.DB)

	userSvc := user.NewService(userRepo, a.Config.JWTSecret)

	userHandler := user.NewHandler(userSvc)

	router.POST("/register", userHandler.Register)

	router.POST("/login",userHandler.Login)

	return router
}
