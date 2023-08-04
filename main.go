package main

import (
	"SimonBK_Login/db"
	"SimonBK_Login/routers"
	"fmt"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	// Establecer la conexión con la base de datos
	err := db.ConnectDB()

	if err != nil {
		fmt.Println("Error al conectar con la base de datos:", err)
		return
	}

	// Configurar e iniciar el enrutador
	r := routers.SetupRouter()

	// Imprimir todas las rutas disponibles
	for _, route := range r.Routes() {
		fmt.Println(route.Path)
	}

	// Configurar la señal de captura
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)

	go func() {
		<-c
		// Código de limpieza: cierra la conexión a la base de datos
		db.CloseDB()
		os.Exit(0)
	}()

	// Escuchar y servir
	err = r.Run(":8000") // escucha y sirve en 0.0.0.0:8000 (por defecto)

	if err != nil {
		fmt.Println("Error al iniciar el servidor:", err)
		return
	}
}
