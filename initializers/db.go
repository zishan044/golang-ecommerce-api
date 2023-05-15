package initializers

import (
	"fmt"
	"os"

	"github.com/golang-ecommerce-api/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Config struct {
	Host     string
	Port     string
	User     string
	Password string
	DBName   string
	SSLMode  string
}

var DB *gorm.DB

func NewConnection(config *Config) {

	var err error

	dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
		config.Host, config.Port, config.User, config.Password, config.DBName, config.SSLMode)

	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Println("could not connect to database")
		os.Exit(-1)
	}
}

func SyncDB() {
	DB.AutoMigrate(&models.Cart{})
	DB.AutoMigrate(&models.Category{})
	DB.AutoMigrate(&models.Customer{})
	DB.AutoMigrate(&models.Order{})
	DB.AutoMigrate(&models.Product{})
	DB.AutoMigrate(&models.Review{})
	DB.AutoMigrate(&models.Supplier{})
}
