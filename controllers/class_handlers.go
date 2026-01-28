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
	migrations.CreateClassInDatabase(className, services.DB, requestCtx)
}