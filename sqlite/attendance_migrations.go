package migrations

import (
	"context"
	"time"

	"github.com/AbanoubGirges/malaykaproject/models"
	"gorm.io/gorm"
)

func CreateClassAttendanceInDatabase(students []models.AttendaceInDatabase, db *gorm.DB, ctx context.Context) error {
	dbCtx, cancel := context.WithTimeout(ctx, time.Second*15)
	defer cancel()
	result:=db.WithContext(dbCtx).Create(&students)
	return result.Error
}
func ReadClassAttendanceFromDatabase(classId uint, date string, db *gorm.DB, ctx context.Context) ([]models.AttendaceInDatabase,error) {
	dbCtx, cancel := context.WithTimeout(ctx, time.Second*15)
	defer cancel()
	var attendance []models.AttendaceInDatabase
	result:=db.WithContext(dbCtx).Where("class_id = ? AND date = ?", classId, date).Find(&attendance)
	return attendance,result.Error
}
func UpdateClassAttendanceInDatabase(attendance models.AttendaceInDatabase, db *gorm.DB, ctx context.Context) error {
	dbCtx, cancel := context.WithTimeout(ctx, time.Second*15)
	defer cancel()
	result := db.WithContext(dbCtx).Model(&models.AttendaceInDatabase{}).Where("class_id = ? AND date = ? AND student = ?", attendance.ClassID, attendance.Date, attendance.Student).Updates(map[string]interface{}{
		"present": attendance.Present,
		"blame": attendance.Blame,
	})
	return result.Error
}
func DeleteClassAttendanceFromDatabase(attendance models.AttendaceInDatabase, db *gorm.DB, ctx context.Context) error {
	dbCtx, cancel := context.WithTimeout(ctx, time.Second*15)
	defer cancel()
	result := db.WithContext(dbCtx).Where("class_id = ? AND date = ? AND student = ?", attendance.ClassID, attendance.Date, attendance.Student).Delete(&models.AttendaceInDatabase{})
	return result.Error
}