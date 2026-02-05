package controllers

import (
	"context"
	"fmt"
	"net/http"
	"time"
	"errors"
	"github.com/AbanoubGirges/malaykaproject/services"
	migrations "github.com/AbanoubGirges/malaykaproject/sqlite"
	//"github.com/AbanoubGirges/malaykaproject/services"
)

func CreateClassHandler(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	requestCtx, cancel := context.WithTimeout(ctx, 15*time.Second)
	defer cancel()
	className := r.URL.Query().Get("class_name")
	err := migrations.CreateClassInDatabase(className, services.DB, requestCtx)
	if err != nil {
		//http.Error(w, "Failed to create class", http.StatusInternalServerError)
		services.RespondWithJson(w, http.StatusInternalServerError, "Failed to create class")
		return
	}
	select {
	case <-requestCtx.Done():
		services.RequestTimeout(w, r)
		return
	default:
		services.RespondWithJson(w, http.StatusCreated, map[string]interface{}{"message": "Class created successfully"})
	}
}
func ReadClassHandler(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	//fmt.Printf("Reading class: %s\n", className)
	readCtx, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel()

	class, err := migrations.ReadClass( services.DB, readCtx)
	if err != nil {
		if errors.Is(err, context.DeadlineExceeded) {
			services.RequestTimeout(w, r)
			return
		}

		services.RespondWithJson(w, 500, map[string]string{
			"error": "FAILED_TO_READ",
		})
		return
	}

	services.RespondWithJson(w, 200, class)
}

func DeleteClassHandler(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	requestCtx, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel()
	className := r.URL.Query().Get("class_name")
	if className == "" {
		services.RespondWithJson(w, http.StatusBadRequest, map[string]string{"error": "MISSING_CLASS_NAME"})
		return
	}
	fmt.Printf("Attempting to delete class: %s\n", className)
	err := migrations.DeleteClassFromDatabase(className, services.DB, requestCtx)
	if err != nil {
		services.RespondWithJson(w, http.StatusInternalServerError, map[string]string{"error": "FAILED_TO_DELETE_CLASS"})
		fmt.Println(err)
		return
	}
	select {
	case <-requestCtx.Done():
		services.RequestTimeout(w, r)
		return
	default:
		services.RespondWithJson(w, http.StatusOK, map[string]interface{}{"message": "Class deleted successfully"})
	}
}
func UpdateClassHandler(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	requestCtx, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel()
	className := r.URL.Query().Get("class_name")
	newName := r.URL.Query().Get("new_name")
	err := migrations.UpdateClassInDatabase(className, newName, services.DB, requestCtx)
	if err != nil {
		services.RespondWithJson(w, http.StatusInternalServerError, map[string]interface{}{"error": "FAILED_TO_UPDATE_CLASS"})
		return
	}
	select {
	case <-requestCtx.Done():
		services.RequestTimeout(w, r)
		return
	default:
		services.RespondWithJson(w, http.StatusOK, map[string]interface{}{"message": "Class updated successfully"})
	}
	//services.RespondWithJson(w, http.StatusNotImplemented, "Update class not implemented yet")
}
