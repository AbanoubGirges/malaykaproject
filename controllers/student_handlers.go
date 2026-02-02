package controllers

import (
	"context"
	"encoding/json"
	"fmt"

	//"errors"
	"net/http"
	"time"

	"github.com/AbanoubGirges/malaykaproject/models"
	"github.com/AbanoubGirges/malaykaproject/services"
	migrations "github.com/AbanoubGirges/malaykaproject/sqlite"
	"github.com/google/uuid"
)

func CreateStudentHandler(w http.ResponseWriter, r *http.Request) {
	claims := r.Context().Value("claims").(map[string]interface{})
	requestCtx, cancel := context.WithTimeout(r.Context(), 15*time.Second)
	defer cancel()
	var student models.Student
	err := json.NewDecoder(r.Body).Decode(&student)
	if err != nil {
		services.RespondWithJson(w, 400, map[string]string{"error": "BAD_REQUEST"})
		fmt.Println(err)
		return
	}
	student.ID = uuid.New().ID()
	if claims["role"].(string) != "admin" {
		student.Class = uint(claims["class"].(float64))
		fmt.Println("student.Class=", student.Class)
	}
	fmt.Println(claims["role"])
	if student.Name == "" || student.Age == 0 || student.Class == 0 || student.Location == "" || student.PhoneNumber == nil || student.ID == 0 || student.Coordinates == "" || student.Birthdate == "" {
		services.RespondWithJson(w, 400, map[string]string{"error": "MISSING_FIELDS"})
		fmt.Println(student.Class)
		return
	}
	err = migrations.CreateStudentInDatabase(services.ToStudentInDatabase(student), services.DB, requestCtx)
	if err != nil {
		services.RespondWithJson(w, 500, map[string]string{"error": "INTERNAL_SERVER_ERROR"})
		return
	}
	select {
	case <-requestCtx.Done():
		services.RespondWithJson(w, 504, map[string]string{"error": "REQUEST_TIMEOUT"})
		return
	default:
		services.RespondWithJson(w, 201, map[string]string{"message": "STUDENT_CREATED"})
	}
}
func ReadStudentHandler(w http.ResponseWriter, r *http.Request) {
	claims := r.Context().Value("claims").(map[string]interface{})
	requestCtx, cancel := context.WithTimeout(r.Context(), 10*time.Second)
	defer cancel()
	if claims["class"] == nil {
		services.RespondWithJson(w, 400, map[string]string{"error": "BAD_REQUEST"})
		return
	}
	students, err := migrations.ReadStudent(uint(claims["class"].(float64)), services.DB, requestCtx)
	if err != nil {
		services.RespondWithJson(w, 500, map[string]string{"error": "INTERNAL_SERVER_ERROR"})
		return
	}
	services.RespondWithJson(w, 200, students)
}
func DeleteStudentHandler(w http.ResponseWriter, r *http.Request) {
	claims := r.Context().Value("claims").(map[string]interface{})
	requestCtx, cancel := context.WithTimeout(r.Context(), 10*time.Second)
	defer cancel()
	err := migrations.DeleteStudentFromDatabase(uint(claims["ID"].(float64)), services.DB, requestCtx)
	if err != nil {
		services.RespondWithJson(w, 500, map[string]string{"error": "INTERNAL_SERVER_ERROR"})
		return
	}
	select {
	case <-requestCtx.Done():
		services.RespondWithJson(w, 504, map[string]string{"error": "REQUEST_TIMEOUT"})
		return
	default:
		services.RespondWithJson(w, 200, map[string]string{"message": "STUDENT_DELETED"})
	}
}
func UpdateStudentHandler(w http.ResponseWriter, r *http.Request) {
	claims := r.Context().Value("claims").(map[string]interface{})
	requestCtx, cancel := context.WithTimeout(r.Context(), 15*time.Second)
	defer cancel()
	var student models.Student
	err := json.NewDecoder(r.Body).Decode(&student)
	if err != nil {
		services.RespondWithJson(w, 400, map[string]string{"error": "BAD_REQUEST"})
		return
	}
	student.Class = uint(claims["class"].(float64))
	err = migrations.UpdateStudentInDatabase(services.ToStudentInDatabase(student), services.DB, requestCtx)
	if err != nil {
		services.RespondWithJson(w, 500, map[string]string{"error": "INTERNAL_SERVER_ERROR"})
		return
	}
	select {
	case <-requestCtx.Done():
		services.RespondWithJson(w, 504, map[string]string{"error": "REQUEST_TIMEOUT"})
		return
	default:
		services.RespondWithJson(w, 200, map[string]string{"message": "STUDENT_UPDATED"})
		return
	}
}
