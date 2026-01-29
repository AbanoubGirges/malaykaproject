package controllers

import (
	"context"
	"net/http"

	"github.com/AbanoubGirges/malaykaproject/services"
	migrations "github.com/AbanoubGirges/malaykaproject/sqlite"
	//"github.com/AbanoubGirges/malaykaproject/services"
)

func CreateClassHandler(w http.ResponseWriter, r *http.Request){
	ctx:= r.Context()
	requestCtx, cancel:= context.WithTimeout(ctx, 15)
	defer cancel()
	className:= r.URL.Query().Get("class_name")
	err:=migrations.CreateClassInDatabase(className, services.DB, requestCtx)
	if err != nil {
		//http.Error(w, "Failed to create class", http.StatusInternalServerError)
		services.RespondWithJson(w, http.StatusInternalServerError, "Failed to create class")
		return
	}
	services.RespondWithJson(w, http.StatusCreated, map[string]interface{}{"message": "Class created successfully"})
}
func ReadClassHandler(w http.ResponseWriter, r *http.Request){
	claims:= r.Context().Value("claims").(map[string]interface{})
	
}