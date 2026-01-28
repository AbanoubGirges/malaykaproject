package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/AbanoubGirges/malaykaproject/routes"
	"github.com/AbanoubGirges/malaykaproject/services"
	migrations "github.com/AbanoubGirges/malaykaproject/sqlite"
	"github.com/joho/godotenv"
)

func main() {
	// ctx:=context.Background()
	var DB = migrations.SetupDatabase()
	services.SetDB(DB)
	godotenv.Load(".env")
	portForServer := os.Getenv("PORT")
	if portForServer == "" {
		log.Println("Failed to get PORT from .env file")
		return
	}
	secretKey := os.Getenv("SECRET_KEY")
	if secretKey == "" {
		log.Println("Failed to get SECRET_KEY from .env file")
		return
	}
	services.TakeSecretKey(secretKey)
	router := routes.SetupRouter(portForServer)
	err := http.ListenAndServe(fmt.Sprintf(":%s", portForServer), router)

	if err != nil {
		log.Fatal("Failed to start server:", err.Error())
	}
}
