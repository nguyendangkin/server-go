package config

import (
	"chin_server/internal/model"
	"fmt"
	"log"

	"github.com/spf13/viper"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// struct để bind data từ yaml vào
type Config struct {
	Server struct {
		Port int
	}
	Database struct {
		Host     string
		Port     int
		User     string
		Password string
		Name     string
	}
}

// nạp config từ yaml vào
func LoadConfig() *Config {
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath("./config") // là file ở trong folder config

	err := viper.ReadInConfig() // đọc file config.yaml
	if err != nil {
		log.Fatalf("error loading config: %v", err)
	}

	var cfg Config
	err = viper.Unmarshal(&cfg) // parse yaml vào struct config trên (bind)
	if err != nil {
		log.Fatalf("error unmarshal config %v", err)
	}
	return &cfg // return con trỏ để tương tác thẳng vào gốc
}

// kết nối vào database qua các biến config (struct)
func ConnectDatabase(cfg *Config) *gorm.DB {
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%d sslmode=disable TimeZone=Asia/Ho_Chi_Minh",
		cfg.Database.Host,
		cfg.Database.User,
		cfg.Database.Password,
		cfg.Database.Name,
		cfg.Database.Port,
	)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(fmt.Sprintf("Failed to connect to database: %v", err))
	}

	fmt.Println("Database connected!")

	// auto-migrate
	if err := db.AutoMigrate(&model.User{}); err != nil {
		panic(fmt.Sprintf("Failed to migration database: %v", err))
	}

	return db
}
