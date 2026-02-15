package controllers

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"time"

	"github.com/AbanoubGirges/malaykaproject/models"
	"github.com/AbanoubGirges/malaykaproject/services"
	migrations "github.com/AbanoubGirges/malaykaproject/sqlite"
	"github.com/google/uuid"
)

// This is working fine
func SignupHandler(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	requestCtx, cancel := context.WithTimeout(ctx, time.Second*15)
	defer cancel()
	var user models.User
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		services.RespondWithJson(w, 400, struct{Error string}{Error: "FAILED_TO_DECODE"})
		return
	}
	if user.Username == "" || user.PhoneNumber == "" || user.Password == "" {
		services.RespondWithJson(w, 400, struct{Error string}{Error: "MISSING_FIELDS"})
		return
	}

	user.ID = uuid.New().ID()
	user.Password, _ = services.HashPassword(user.Password)
	userInDatabase := services.ToUserInDatabase(user)
	err := migrations.CreateUserInDatabase(userInDatabase, services.DB, requestCtx)
	if err != nil {
		services.RespondWithJson(w, 500, struct{Error string}{Error: "FAILED_TO_CREATE_USER"})
		return
	}

	select {
	case <-requestCtx.Done():
		services.RequestTimeout(w, r)
		return
	default:
		services.RespondWithJson(w, 201, map[string]interface{}{"message": "User created successfully"})
	}
}
func LoginHandler(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	requestCtx, cancel := context.WithTimeout(ctx, time.Second*10)
	defer cancel()
	var userLogin models.UserLoginRequest
	if err := json.NewDecoder(r.Body).Decode(&userLogin); err != nil {
		services.RespondWithJson(w, 400, struct{Error string}{Error: "FAILED_TO_DECODE"})
		return
	}
	//defer services.RespondWithJson(w, 200, struct{}{})
	userInDatabase, err := migrations.FetchUserLogin(userLogin.PhoneNumber, services.DB, requestCtx, userLogin.Password)
	if err != nil {
		services.RespondWithJson(w, 401, struct{}{})
		return
	}
	if errors.Is(err, errors.New("user not found")) {
		services.RespondWithJson(w, 404, struct{Error string}{Error: "USER_NOT_FOUND"})
		return
	}
	err = services.CheckPasswordHash(userLogin.Password, userInDatabase.Password)
	if err != nil {
		services.RespondWithJson(w, 401, struct{Error string}{Error: "INVALID_PASSWORD"})
		return
	}
	tokenString, err := services.GenerateJWT(userInDatabase, services.SecretKey)
	if err != nil {
		services.RespondWithJson(w, 500, struct{Error string}{Error: "FAILED_TO_GENERATE_TOKEN"})
		return
	}
	select {
	case <-requestCtx.Done():
		services.RequestTimeout(w, r)
		return
	default:
		services.RespondWithJson(w, 200, struct{ Token string 
			Name string
			 Role string 
			 Class uint }{Token: tokenString, Name: userInDatabase.Name, Role: userInDatabase.Role, Class: userInDatabase.Class})
		return
	}
	
}
func EditProfileHandler(w http.ResponseWriter, r *http.Request) {
	claims:= r.Context().Value("claims").(map[string]interface{})
	requestCtx, cancel:= context.WithTimeout(r.Context(), 15*time.Second)
	defer cancel()
	editField:=r.URL.Query().Get("field")
	fmt.Printf("field= %v",editField)
	newValue:=r.URL.Query().Get("value")
	fmt.Printf("value= %v",newValue)
	if editField=="" || newValue==""{
		services.RespondWithJson(w,400,struct{Error string}{Error:"MISSING_FIELDS"})
		return
	}
	err:=migrations.UpdateUserInDatabaseField(uint(claims["user_id"].(float64)), editField, newValue, services.DB, requestCtx)
	if err!=nil{
		services.RespondWithJson(w,500,struct{Error string}{Error:"INTERNAL_SERVER_ERROR"})
		return
	}
}
func DeleteUserHandler(w http.ResponseWriter, r *http.Request) {
	claims:= r.Context().Value("claims").(map[string]interface{})
	requestCtx, cancel:= context.WithTimeout(r.Context(), 15*time.Second)
	defer cancel()

	err := migrations.DeleteUserFromDatabase(uint(claims["user_id"].(float64)), services.DB, requestCtx)
	if err != nil {
		services.RespondWithJson(w, 500, struct{Error string}{Error:"INTERNAL_SERVER_ERROR"})
		return
	}
	services.RespondWithJson(w, 200, struct{Message string}{Message:"User deleted successfully"})
}