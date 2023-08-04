package main

import (
	"SimonBK_Login/db"
	"SimonBK_Login/models"
	"fmt"
	"log"
)

func main() {
	err := db.ConnectDB()
	if err != nil {
		log.Fatalf("Could not connect to DB: %v", err)
	}

	// Asegúrate de definir el orden correcto de creación de las tablas
	// según sus relaciones
	err = db.DBConn.AutoMigrate(&models.Compania{}) // Primero, la tabla "compania"
	if err != nil {
		log.Fatalf("Failed to migrate table Compania: %v", err)
	}

	err = db.DBConn.AutoMigrate(&models.Usuario{}) // Luego, la tabla "usuario"
	if err != nil {
		log.Fatalf("Failed to migrate table Usuario: %v", err)
	}

	err = db.DBConn.AutoMigrate(&models.Cliente{}) // Luego, la tabla "cliente"
	if err != nil {
		log.Fatalf("Failed to migrate table Cliente: %v", err)
	}

	err = db.DBConn.AutoMigrate(&models.Modulo{}) // Luego, la tabla "modulo"
	if err != nil {
		log.Fatalf("Failed to migrate table Modulo: %v", err)
	}

	err = db.DBConn.AutoMigrate(&models.Roles{}) // Luego, la tabla "rol"
	if err != nil {
		log.Fatalf("Failed to migrate table Rol: %v", err)
	}

	err = db.DBConn.AutoMigrate(&models.PermisoUsuario{}) // Finalmente, la tabla "permiso_usuario"
	if err != nil {
		log.Fatalf("Failed to migrate table PermisoUsuario: %v", err)
	}

	fmt.Println("Migration successful")
}
