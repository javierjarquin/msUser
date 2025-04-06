package main

import (
	"log"
	"msUser/config"
	"msUser/internal/delivery/http"
	"msUser/internal/delivery/router"
	"msUser/internal/repository"
	usercase "msUser/internal/usecase"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	// Cargar el archivo .env
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

	// Crear el router de Gin
	r := gin.Default()

	// Configurar CORS dinámico usando gin-contrib/cors
	corsConfig := cors.Config{
		AllowOrigins: []string{
			"http://localhost:33313", // Puerto de Flutter (ajustado a tu configuración)
			"http://localhost:4200",  // Otro puerto común para aplicaciones frontend (Angular, por ejemplo)
			"http://localhost:8080",  // Si tienes otros orígenes locales en tu máquina, agrégales aquí
		},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Content-Type", "Authorization"},
		AllowCredentials: true,
	}

	// Aplicar CORS usando el middleware de gin-contrib/cors
	r.Use(cors.New(corsConfig))

	// Middleware para registrar todas las solicitudes excepto OPTIONS
	r.Use(func(c *gin.Context) {
		if c.Request.Method != "OPTIONS" {
			log.Printf("Request: %s %s", c.Request.Method, c.Request.URL.Path)
		}
		c.Next()
	})

	// Registrar rutas
	router.NewRouter(userHandler).RegisterRoutes(r)

	// Iniciar servidor
	log.Println("Servidor corriendo en el puerto", cfg.Port)
	if err := r.Run(":" + cfg.Port); err != nil {
		log.Fatal("Error al iniciar el servidor:", err)
	}
}
