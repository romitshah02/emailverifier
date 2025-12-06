package main

import (
	"fmt"
	"gateway/internal/db"
	"gateway/internal/gateway"
	"gateway/internal/routes"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {

	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found or failed to load")
	}

	db.Init()
	r := gin.Default()
	routes.RegisterRoutes(r)

	proxy := r.Group("/proxy")
	proxy.Any("/*path", gateway.ProxyHandler)
	gateway.StartRouteWatcher()

	r.Run(":8080")

	fmt.Printf("Running on port 8080")
}
