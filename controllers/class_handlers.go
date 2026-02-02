package controllers

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/AbanoubGirges/malaykaproject/services"
	migrations "github.com/AbanoubGirges/malaykaproject/sqlite"
	//"github.com/AbanoubGirges/malaykaproject/services"
)

func CreateClassHandler(w http.ResponseWriter, r *http.Request){
	ctx:= r.Context()
	requestCtx, cancel:= context.WithTimeout(ctx, 15*time.Second)
	defer cancel()
	className:= r.URL.Query().Get("class_name")
	err:=migrations.CreateClassInDatabase(className, services.DB, requestCtx)
	if err != nil {
		//http.Error(w, "Failed to create class", http.StatusInternalServerError)
		services.RespondWithJson(w, http.StatusInternalServerError, "Failed to create class")
		return
	}
	select{
	case<-requestCtx.Done():
		services.RequestTimeout(w,r)
		return
	default:
		services.RespondWithJson(w, http.StatusCreated, map[string]interface{}{"message": "Class created successfully"})
	}
}
func ReadClassHandler(w http.ResponseWriter, r *http.Request){
	claims:= r.Context().Value("claims").(map[string]string)
	readCtx,cancel:=context.WithTimeout(r.Context(),time.Second*10)
	defer cancel()
	class,err:=migrations.ReadClass(claims["class"],services.DB,readCtx)
	if err!=nil{
		services.RespondWithJson(w,500,struct{error string}{error:"FAILED_TO_READ"})
	}
	select{
	case<-readCtx.Done():
		services.RequestTimeout(w,r)
		return
	default:
		services.RespondWithJson(w,200,class)
	}
}
func DeleteClassHandler(w http.ResponseWriter,r *http.Request){
	ctx:= r.Context()
	requestCtx, cancel:= context.WithTimeout(ctx, 10*time.Second)
	defer cancel()
	className:= r.URL.Query().Get("class_name")
	if className==""{
		services.RespondWithJson(w, http.StatusBadRequest, map[string]string{"error": "MISSING_CLASS_NAME"})
		return
	}
	err:=migrations.DeleteClassFromDatabase(className, services.DB, requestCtx)
	if err != nil {
		services.RespondWithJson(w, http.StatusInternalServerError, map[string]string{"error": "FAILED_TO_DELETE_CLASS"})
		fmt.Println(err)
		return
	}
	select{
	case<-requestCtx.Done():
		services.RequestTimeout(w,r)
		return
	default:
	services.RespondWithJson(w, http.StatusOK, map[string]interface{}{"message": "Class deleted successfully"})
	}
}
func UpdateClassHandler(w http.ResponseWriter,r *http.Request){
	ctx:=r.Context()
	requestCtx,cancel:=context.WithTimeout(ctx,10*time.Second)
	defer cancel()
	className:=r.URL.Query().Get("class_name")
	newName:=r.URL.Query().Get("new_name")
	err:=migrations.UpdateClassInDatabase(className,newName,services.DB,requestCtx)
	if err!=nil{
		services.RespondWithJson(w,http.StatusInternalServerError,map[string]interface{}{"error": "FAILED_TO_UPDATE_CLASS"})
		return
	}
	select{
	case<-requestCtx.Done():
		services.RequestTimeout(w,r)
		return
	default:
		services.RespondWithJson(w,http.StatusOK,map[string]interface{}{"message":"Class updated successfully"})
	}
	//services.RespondWithJson(w, http.StatusNotImplemented, "Update class not implemented yet")
}
