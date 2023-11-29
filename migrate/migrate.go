package migrate

// func RunMigrations() {
// 	err := db.ConnectDB()
// 	if err != nil {
// 		log.Fatalf("Could not connect to DB: %v", err)
// 	}

// Asegúrate de definir el orden correcto de creación de las tablas
// según sus relaciones
// err = db.DBConn.AutoMigrate(&models.Company{}) // Primero, la tabla "compania"
// if err != nil {
// 	log.Fatalf("Failed to migrate table Compania: %v", err)
// }

// err = db.DBConn.AutoMigrate(&models.User{}) // Luego, la tabla "usuario"
// if err != nil {
// 	log.Fatalf("Failed to migrate table Usuario: %v", err)
// }

// err = db.DBConn.AutoMigrate(&models.Customer{}) // Luego, la tabla "cliente"
// if err != nil {
// 	log.Fatalf("Failed to migrate table Cliente: %v", err)
// }

// err = db.DBConn.AutoMigrate(&models.Module{}) // Luego, la tabla "modulo"
// if err != nil {
// 	log.Fatalf("Failed to migrate table Modulo: %v", err)
// }

// err = db.DBConn.AutoMigrate(&models.Role{}) // Luego, la tabla "rol"
// if err != nil {
// 	log.Fatalf("Failed to migrate table Rol: %v", err)
// }

// err = db.DBConn.AutoMigrate(&models.UserPermission{}) // Finalmente, la tabla "permiso_usuario"
// if err != nil {
// 	log.Fatalf("Failed to migrate table PermisoUsuario: %v", err)
// }
// 	err = db.DBConn.AutoMigrate(&models.RefreshToken{}) // Finalmente, la tabla "permiso_usuario"
// 	if err != nil {
// 		log.Fatalf("Failed to migrate table PermisoUsuario: %v", err)
// 	}

// 	fmt.Println("Migration successful")
// }
