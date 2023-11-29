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
	usuario := r.Group("/users")
	{
		usuario.POST("/login/", controllers.Login)
		// usuario.GET("/resources", controllers.Resources) //login
		// usuario.POST("/refresh", controllers.Refresh)    // Refresh
	}
	return r
}
