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
