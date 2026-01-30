package routes

import (
	"net/http"

	"github.com/AbanoubGirges/malaykaproject/controllers"
	custommiddleware "github.com/AbanoubGirges/malaykaproject/middleware"
	"github.com/AbanoubGirges/malaykaproject/services"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
)

func SetupRouter(portForServer string) *chi.Mux {
	Router := chi.NewRouter()
	Router.Use(middleware.Logger)
	Router.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"*"},
		ExposedHeaders:   []string{"*"},
		AllowCredentials: true,
	}))
	println("Server will start at port:", portForServer)
	Router.Get("/ready", func(w http.ResponseWriter, r *http.Request) { services.RespondWithJson(w, 200, struct{}{}) })
	Router.Post("/signup", controllers.SignupHandler)
	Router.Get("/login", controllers.LoginHandler)
	//class routes
	classRouter:=chi.NewRouter()
	classRouter.Post("/create", custommiddleware.AdminAuthMiddleware(http.HandlerFunc(controllers.CreateClassHandler)).ServeHTTP)
	classRouter.Get("/read",custommiddleware.UserAuthMiddleware(http.HandlerFunc(controllers.ReadClassHandler)).ServeHTTP)
	classRouter.Delete("/delete",custommiddleware.AdminAuthMiddleware(http.HandlerFunc(controllers.DeleteClassHandler)))
	classRouter.Put("/update",custommiddleware.AdminAuthMiddleware(http.HandlerFunc(controllers.UpdateClassHandler)))
	Router.Mount("/class", classRouter)
	//student routes
	studentRouter:=chi.NewRouter()
	studentRouter.Post("/create",custommiddleware.UserAuthMiddleware(http.HandlerFunc(controllers.CreateStudentHandler)).ServeHTTP)
	studentRouter.Get("/read",custommiddleware.UserAuthMiddleware(http.HandlerFunc(controllers.ReadStudentHandler)).ServeHTTP)
	studentRouter.Delete("/delete",custommiddleware.UserAuthMiddleware(http.HandlerFunc(controllers.DeleteStudentHandler)).ServeHTTP)
	studentRouter.Put("/update",custommiddleware.UserAuthMiddleware(http.HandlerFunc(controllers.UpdateStudentHandler)).ServeHTTP)
	Router.Mount("/student", studentRouter)
	return Router
}
