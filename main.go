package main

import (
	"log"
	"msUser/config"
	"msUser/internal/delivery/http"
	"msUser/internal/delivery/router"
	"msUser/internal/repository"
	usercase "msUser/internal/usecase"

	"github.com/joho/godotenv"
)

func main() {

	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file")
	}
	// Cargar configuración desde variables de entorno
	cfg := config.LoadConfig()

	// Conectar a la base de datos
	userRepo, err := repository.OpenUserDb(cfg.DBUrl)
	if err != nil {
		log.Fatal("Error al conectar a la BD:", err)
	}

	// Iniciar dependencias
	userUseCase := usercase.NewUserCase(userRepo)
	userHandler := http.NewUserHandler(userUseCase)

	// Configurar rutas
	r := router.NewRouter(userHandler)

	// Iniciar servidor
	log.Println("Servidor corriendo en el puerto", cfg.Port)
	r.Run(":" + cfg.Port)
}
