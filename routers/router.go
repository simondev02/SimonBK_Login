package routers

import (
	"SimonBK_Login/controllers"

	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()
	// Rutas para el modelo Usuario
	usuario := r.Group("/usuario")
	{
		usuario.POST("/login", controllers.Login) //login
	}
	return r
}
