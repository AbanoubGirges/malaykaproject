package controllers

import (
	"context"
	"encoding/json"
	"net/http"
	"time"
	"errors"
	"github.com/AbanoubGirges/malaykaproject/models"
	"github.com/AbanoubGirges/malaykaproject/services"
	migrations "github.com/AbanoubGirges/malaykaproject/sqlite"
	"github.com/google/uuid"
)

// var userPtr *models.User
func SignupHandler(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	requestCtx, cancel := context.WithTimeout(ctx, time.Second*10)
	defer cancel()
	var user models.User
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		services.RespondWithJson(w, 400, struct{}{})
		return
	}
	if user.Username == "" || user.PhoneNumber == "" || user.Password == "" {
		services.RespondWithJson(w, 400, struct{}{})
		return
	}
	defer services.RespondWithJson(w, 201, struct{}{})
	user.ID = uuid.New().ID()
	user.Password, _ = services.HashPassword(user.Password)
	userInDatabase := services.ToUserInDatabase(user)
	err := migrations.CreateUserInDatabase(userInDatabase, services.DB, requestCtx)
	if err != nil {
		services.RespondWithJson(w, 500, struct{}{})
		return
	}
	select {
	case <-requestCtx.Done():
		services.RequestTimeout(w, r)
		return

	}
}
func LoginHandler(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	requestCtx, cancel := context.WithTimeout(ctx, time.Second*10)
	defer cancel()
	var userLogin models.UserLoginRequest
	if err := json.NewDecoder(r.Body).Decode(&userLogin); err != nil {
		services.RespondWithJson(w, 400, struct{}{})
		return
	}
	defer services.RespondWithJson(w, 200, struct{}{})
	userInDatabase, err := migrations.FetchUserLogin(userLogin.PhoneNumber, services.DB, requestCtx, userLogin.Password)
	if err != nil {
		services.RespondWithJson(w, 401, struct{}{})
		return
	}
	if errors.Is(err, errors.New("user not found")) {
		services.RespondWithJson(w, 404, struct{}{})
		return
	}
	err = services.CheckPasswordHash(userLogin.Password, userInDatabase.Password)
	if err != nil {
		services.RespondWithJson(w, 401, struct{}{})
		return
	}
	select {
	case <-requestCtx.Done():
		services.RequestTimeout(w, r)
		return

	}
}
