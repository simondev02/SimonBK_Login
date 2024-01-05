package routers

import (
	"SimonBK_Login/api/controllers"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()

	// Configurar CORS
	config := cors.DefaultConfig()
	config.AllowAllOrigins = true
	r.Use(cors.New(config))

	// Rutas para el modelo Usuario
	auth := r.Group("/auth")
	{
		auth.POST("/login/", controllers.Login)
		auth.GET("/resources", controllers.Resources)          //login
		auth.POST("/refresh", controllers.RefreshTokenHandler) // Refresh
	}
	return r
}
