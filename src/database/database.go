package database

import (
	"errors"
	"fmt"
	"golang-api/src/helpers"
	"log"

	"github.com/spf13/viper"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

// InitDatabase inicializa la conexión a la base de datos
func InitDatabase() {
	var err error
	host := helpers.Coalesce(viper.GetString("DB_HOST"), "localhost")
	user := helpers.Coalesce(viper.GetString("DB_USER"), "root")
	pass := helpers.Coalesce(viper.GetString("DB_PASS"), "root")
	port := helpers.Coalesce(viper.GetInt("DB_PORT"), 3306)
	name := helpers.Coalesce(viper.GetString("DB_NAME"), "golang_api")

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local", user, pass, host, port, name)

	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("no se pudo conectar a la base de datos:", err)
	}

	log.Println("conexión a la base de datos MySQL exitosa")
}

func CheckConnection(db *gorm.DB) error {
	connection, err := db.DB()
	if err != nil {
		return errors.New("no se pudo obtener la conexión subyacente de GORM")
	}

	if err := connection.Ping(); err != nil {
		return errors.New("no se pudo conectar a la base de datos: " + err.Error())
	}

	return nil
}
