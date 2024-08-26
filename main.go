package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"os"
	"user-management-backend/internal/controller"
	"user-management-backend/internal/model"
	"user-management-backend/internal/repository"
	"user-management-backend/internal/service"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type DBConfig struct {
	Username  string `json:"username"`
	Password  string `json:"password"`
	Host      string `json:"host"`
	Port      int    `json:"port"`
	DBName    string `json:"dbname"`
	Charset   string `json:"charset"`
	ParseTime bool   `json:"parseTime"`
	Loc       string `json:"loc"`
}

// LoadDBConfig loads the database configuration from a db.json config file
func LoadDBConfig() (*DBConfig, error) {
	file, err := os.Open("./config/db.json")
	if err != nil {
		return nil, err
	}
	defer file.Close()

	dbconfig, err := io.ReadAll(file)
	if err != nil {
		return nil, err
	}

	var config DBConfig
	if err := json.Unmarshal(dbconfig, &config); err != nil {
		return nil, err
	}

	return &config, nil
}

func main() {
	config, err := LoadDBConfig()
	if err != nil {
		log.Fatal("Failed to load database configuration:", err)
	}

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=%s&parseTime=%t&loc=%s",
		config.Username, config.Password, config.Host, config.Port,
		config.DBName, config.Charset, config.ParseTime, config.Loc)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Failed to connect to the database: %v", err)
	}

	// Auto-migrate to keep the schema up to date
	db.AutoMigrate(&model.User{})

	// Initialize repository and service
	userRepo := repository.NewUserRepository(db)
	userService := service.NewUserService(userRepo)

	// Initialize Echo instance
	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	// CORS middleware - to allow cross origin requests
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"http://localhost:4200"}, // Angular frontend's URL
		AllowMethods: []string{echo.GET, echo.POST, echo.PUT, echo.DELETE},
	}))
	// Define routes
	controller.NewUserController(e, userService)

	// Start the server
	e.Start(":8080")
}
