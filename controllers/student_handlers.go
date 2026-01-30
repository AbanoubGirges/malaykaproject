package controllers
import (
	"context"
	"encoding/json"
	//"errors"
	"net/http"
	"time"

	"github.com/AbanoubGirges/malaykaproject/models"
	"github.com/AbanoubGirges/malaykaproject/services"
	migrations "github.com/AbanoubGirges/malaykaproject/sqlite"
	"github.com/google/uuid"
)
func CreateStudentHandler(w http.ResponseWriter, r *http.Request){
	ctx:= r.Context()
	requestCtx, cancel:= context.WithTimeout(ctx, 15*time.Second)
	defer cancel()
	_,err:=services.ValidateJWT(w.Header().Get("Authentication"), services.SecretKey)	
	if err!=nil{
		services.RespondWithJson(w,401,struct{error string}{error:"UNAUTHENTICATED"})
	}
	var student models.Student
	err=json.NewDecoder(r.Body).Decode(&student)
	if err!=nil{
		services.RespondWithJson(w,400,struct{error string}{error:"BAD_REQUEST"})
		return
	}
	student.ID=uuid.New().ID()
	if student.Name=="" || student.Age==0|| student.Class==0|| student.Location==""|| student.PhoneNumber==nil||student.ID==0||student.Coordinates==""||student.Birthdate==""{
		services.RespondWithJson(w,400,struct{error string}{error:"MISSING_FIELDS"})
		return
	}
	err=migrations.CreateStudentInDatabase(services.ToStudentInDatabase(student), services.DB, requestCtx)
	if err!=nil{
		services.RespondWithJson(w,500,struct{error string}{error:"INTERNAL_SERVER_ERROR"})
		return
	}
	select{
	case <-requestCtx.Done():
		services.RespondWithJson(w,504,struct{error string}{error:"REQUEST_TIMEOUT"})
		return
	default:
		services.RespondWithJson(w,201,struct{message string}{message:"STUDENT_CREATED"})
	}
}
func ReadStudentHandler(w http.ResponseWriter, r *http.Request){
	ctx:= r.Context()
	requestCtx, cancel:= context.WithTimeout(ctx, 10*time.Second)
	defer cancel()
	 claims,err:=services.ValidateJWT(w.Header().Get("Authentication"), services.SecretKey)
	if err!=nil{
		services.RespondWithJson(w,401,struct{error string}{error:"UNAUTHENTICATED"})
		return
	}
	var students []models.StudentInDatabase
	result:=migrations.ReadStudent(claims["ID"].(uint), services.DB, requestCtx)
	if result.Error!=nil{
		services.RespondWithJson(w,500,struct{error string}{error:"INTERNAL_SERVER_ERROR"})
		return
	}
	services.RespondWithJson(w,200,students)
}
func DeleteStudentHandler(w http.ResponseWriter, r *http.Request){}
func UpdateStudentHandler(w http.ResponseWriter, r *http.Request){}