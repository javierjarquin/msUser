// --
// Manejo de configuraciones
// --
package config

import (
	"log"
	"os"
)

// --
// Estructura de configuraciones
// --
type Config struct {
	Port  string
	DBUrl string
}

// --
// El *Config hace referencia a que vamos a usar la estructura que definimos anteriormente
// --
func LoadConfig() *Config {
	//	--
	//	Asignamos el puerto con la librería os
	//	--
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080" // El puerto por defaul si no hay config es el 8080
	}

	//	--
	//	Asignamos el string de conección para mysql
	//	--
	dbUrl := os.Getenv("DATABASE_URL")
	if dbUrl == "" {
		log.Fatal("Favor de asignar el DATABASE_URL")
	}

	//	--
	//	Asignamos un puntero de memoria & y creamos la intancia de la estructra Config
	//	--
	return &Config{
		Port:  port,
		DBUrl: dbUrl,
	}
}
