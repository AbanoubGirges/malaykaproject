package migrations

import (
	"context"
	"errors"
	"time"

	"github.com/AbanoubGirges/malaykaproject/models"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func SetupDatabase() *gorm.DB {
	db, err := gorm.Open(sqlite.Open("malayka.db"), &gorm.Config{})
	if err != nil {
		panic("Failed to connect database")
	}
	//ctx:=context.Background()
	db.AutoMigrate(&models.UserInDatabase{})

	return db
}
func CreateUserInDatabase(user models.UserInDatabase, db *gorm.DB, ctx context.Context) error {
	dbCtx, cancel := context.WithTimeout(ctx, time.Second*3)
	defer cancel()
	result := db.WithContext(dbCtx).Create(&user)
	return result.Error
}
func FetchUserLogin(phoneNumber string, db *gorm.DB, ctx context.Context, password string) (models.UserInDatabase, error) {

	dbCtx, cancel := context.WithTimeout(ctx, time.Second*10)
	defer cancel()
	var user models.UserInDatabase
	result := db.WithContext(dbCtx).Where("phone_number = ?", phoneNumber).First(&user)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return models.UserInDatabase{}, errors.New("user not found")
	}
	if result.Error != nil {
		return models.UserInDatabase{}, result.Error
	}
	return user, nil
}
