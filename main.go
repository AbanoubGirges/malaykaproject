package main

import (
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
	router := routes.SetupRouter(portForServer)
	http.ListenAndServe(":"+portForServer, router)
}
