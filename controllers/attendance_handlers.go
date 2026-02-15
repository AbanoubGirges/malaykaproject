package controllers

import (
	"context"
	"net/http"
	"strconv"
	"time"
	"encoding/json"
	"github.com/AbanoubGirges/malaykaproject/models"
	"github.com/AbanoubGirges/malaykaproject/services"
	migrations "github.com/AbanoubGirges/malaykaproject/sqlite"
	"github.com/go-chi/chi/v5"
)

func CreateClassAttendance(w http.ResponseWriter, r *http.Request) {
	claims := r.Context().Value("claims").(map[string]interface{})
	requestCtx, cancel := context.WithTimeout(r.Context(), 15*time.Second)
	defer cancel()
	var studentsAttendance []models.AttendaceInDatabase
	var classId uint
	if err := json.NewDecoder(r.Body).Decode(&studentsAttendance); err != nil {
		services.RespondWithJson(w, 400, struct{Error string}{Error: "FAILED_TO_DECODE"})
		return
	}
	for i:= range studentsAttendance{
		studentsAttendance[i].Blame=claims["userId"].(uint32)
		studentsAttendance[i].ClassID=uint32(classId)
	}
	if claims["role"] != "admin" {
		classId = claims["class"].(uint)
	} else {
		id := chi.URLParam(r, "id")
		id64, err := strconv.ParseUint(id, 10, 0)
		if err!=nil{
			services.RespondWithJson(w,500,map[string]string{"error":"FAILED_TO_PARSE_CLASS_ID"})
		}
		classId= uint(id64)
	}
	err:=migrations.CreateClassAttendanceInDatabase(studentsAttendance,services.DB,requestCtx)
	if err!=nil{
		services.RespondWithJson(w,500,map[string]string{"error":"FAILED_TO_CREATE_ATTENDANCE"})
		return
	}
	select {
	case <-requestCtx.Done():
		services.RespondWithJson(w, 504, map[string]string{"error": "REQUEST_TIMEOUT"})
		return
	default:
		services.RespondWithJson(w, 201, map[string]string{"message": "ATTENDANCE_CREATED"})
	}
}
