package main

import (
	"os"
	"github.com/joho/godotenv"
	"log"
)

func main(){
godotenv.Load(".env")
portForServer:=os.Getenv("PORT")
if portForServer==""{ 
	log.Println("Failed to get PORT from .env file")
	return
}
}