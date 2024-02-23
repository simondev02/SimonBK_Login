package routers

import (
	"SimonBK_Login/api/controllers"

	"github.com/gin-gonic/gin"
)

func SetupRouter(r *gin.Engine) {
	// Configurar CORS
	/* config := cors.DefaultConfig()
	config.AllowAllOrigins = true
	r.Use(cors.New(config)) */

	// Rutas para el modelo Usuario
	auth := r.Group("/auth")
	{
		auth.POST("/login/", controllers.Login)
		auth.GET("/resources", controllers.Resources)          //login
		auth.POST("/refresh", controllers.RefreshTokenHandler) // Refresh
	}
}
