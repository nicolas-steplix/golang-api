package services

import (
	"golang-api/src/database"
	"golang-api/src/helpers"

	"github.com/spf13/viper"
)

type HealthService struct{}

func NewHealthService() *HealthService {
	return &HealthService{}
}

func (s *HealthService) Check() (map[string]interface{}, error) {
	env := helpers.Coalesce(viper.GetString("APP_ENV"), "development")
	name := helpers.Coalesce(viper.GetString("NAME"), "GoLang API")
	version := helpers.Coalesce(viper.GetString("VERSION"), "0.0.1")

	// Verificar la conexión a la base de datos
	err := database.CheckConnection(database.DB)
	status := "up"
	if err != nil {
		status = "down"
	}

	// Construir la respuesta
	check := map[string]interface{}{
		"name":     name,
		"version":  version,
		"env":      env,
		"alive":    status == "up",
		"database": status,
	}

	return check, nil
}
