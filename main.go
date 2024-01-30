package main

import (
	"SimonBK_Login/docs"
	"SimonBK_Login/infra/db"
	"SimonBK_Login/routers"
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func main() {
	// Cargar variables de entorno
	err := godotenv.Load()
	if err != nil {
		return
	}
	// Establecer la conexión con la base de datos
	// migrate.RunMigrations() -> ejecutar migraciones, descomentar para crear bbdd
	err = db.ConnectDB()

	// Configurar Swagger
	docs.SwaggerInfo.Title = "Auth API"
	docs.SwaggerInfo.Description = "API for user auth"
	docs.SwaggerInfo.Version = "1.0"
	docs.SwaggerInfo.Host = os.Getenv("SERVICE_HOST") + ":" + os.Getenv("SERVICE_PORT")
	docs.SwaggerInfo.BasePath = "/"

	if err != nil {
		fmt.Println("Error al conectar con la base de datos:", err)
		return
	}
	// Configurar CORS

	r := gin.Default()

	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"POST", "GET", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Accept", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
	}))

	// Configurar e iniciar el enrutador
	r = routers.SetupRouter()

	// Agregar la ruta de Swagger sin el middleware de validación de token
	r.GET("/Login/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	// Imprimir todas las rutas disponibles
	// for _, route := range r.Routes() {
	// 	fmt.Println(route.Path)
	// }

	// Configurar la señal de captura
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)

	go func() {
		<-c
		// Código de limpieza: cierra la conexión a la base de datos
		db.CloseDB()
		os.Exit(0)
	}()

	// // Escuchar y servir
	certFile := os.Getenv("TLS_CERT")
	certKey := os.Getenv("TLS_KEY")
	if certFile == "" || certKey == "" {
		log.Println("Error al leer las variables de entorno.")
		db.CloseDB()
		os.Exit(1)
	}

	err = r.RunTLS(":"+os.Getenv("SERVICE_PORT"), certFile, certKey) // escucha y sirve en 0.0.0.0:60000 (por defecto)
	// err = r.Run(":" + os.Getenv("SERVICE_PORT"))
	if err != nil {
		fmt.Println("Error al iniciar el servidor:", err)
		return
	}
}
