package routes

import (
	
	"net/http"

	"github.com/AbanoubGirges/malaykaproject/controllers"
	"github.com/AbanoubGirges/malaykaproject/services"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
)
func SetupRouter(portForServer string) *chi.Mux{
	Router:=chi.NewRouter()
	Router.Use(middleware.Logger)
	Router.Use(cors.Handler(cors.Options{
		AllowedOrigins:     []string{"*"},
		AllowedMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:     []string{"*"},
		ExposedHeaders:    []string{"*"},
		AllowCredentials: true,
	}))
	println("Server will start at port:", portForServer)
	Router.Get("/ready",func(w http.ResponseWriter, r *http.Request) {services.RespondWithJson(w,200,struct{}{})})
	Router.Post("/signup",controllers.SignupHandler)	
	Router.Get("/login",controllers.LoginHandler)	
	return Router
}
