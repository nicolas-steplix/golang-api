package main

import (
	"golang-api/src/database"
	"golang-api/src/server"
	"log"

	"github.com/spf13/viper"
)

func main() {
	viper.SetConfigFile(".env")
	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("Error al leer el archivo de configuración: %v", err)
	}

	// Inicializar base de datos
	database.InitDatabase()

	// Iniciar el servidor
	if err := server.StartServer(); err != nil {
		log.Fatal("Error al iniciar el servidor:", err)
	}
}
