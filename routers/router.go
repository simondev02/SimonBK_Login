package routers

import (
	"SimonBK_Login/controllers"

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
	usuario := r.Group("/usuario")
	{
		usuario.POST("/login/", controllers.Login) //login
	}

	return r
}
