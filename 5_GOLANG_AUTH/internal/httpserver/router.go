package httpserver

import (
	"go-auth/internal/app"
	"go-auth/internal/middleware"
	"go-auth/internal/user"
	"net/http"

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

	api := router.Group("/api")

	api.Use(middleware.AuthRequired(a.Config.JWTSecret))

	api.GET("/files", func(c *gin.Context){
		userId, _ := middleware.GetUserID(c)
		role, _ := middleware.GetRole(c)

		c.JSON(http.StatusOK, gin.H{
			"ok": true,
			"userId": userId,
			"role": role,
			"files": []any{},
		})
		return
	})

	api.GET("/products", func(c *gin.Context){
		c.JSON(http.StatusOK, gin.H{
			"ok": true,
			"products": []any{},
		})
		return
	})

	admin := api.Group("/admin")

	admin.Use(middleware.RequiredAdmin())

	admin.GET("/getme",func(c *gin.Context){
		role, _ := middleware.GetRole(c)
		c.JSON(http.StatusOK, gin.H{
			"ok": true,
			"role": role,
		})
		return
	})

	return router
}
