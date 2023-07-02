package main

import (
	"blog-app/configs"
	"blog-app/routes"
	"os"

	"github.com/labstack/echo/v4"
)

func main() {
	// loadEnv()
	configs.InitDB()

	e := echo.New()
	e = routes.Router(e)
	e.Start(":" + getPort())

}

func getPort() string {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	return port
}

// func loadEnv() {
// 	err := godotenv.Load()
// 	if err != nil {
// 		log.Fatal("Error loading .env file")
// 	}
// }
