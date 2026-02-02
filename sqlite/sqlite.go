package migrations

import (
	"context"
	"errors"
	"strconv"
	"time"

	"github.com/AbanoubGirges/malaykaproject/models"
	"github.com/google/uuid"
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
	db.AutoMigrate(&models.ClassInDatabase{})
	db.AutoMigrate(&models.StudentInDatabase{})
	return db
}
//user migrations
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
func UpdateUserInDatabaseField(id uint, field string, value string, db *gorm.DB, ctx context.Context) error {
	dbCtx, cancel := context.WithTimeout(ctx, time.Second*5)
	defer cancel()
	result := db.WithContext(dbCtx).Model(&models.UserInDatabase{}).Where("id = ?", id).Update(field, value)
	return result.Error
}
func DeleteUserFromDatabase(id uint, db *gorm.DB, ctx context.Context) error {
	dbCtx, cancel := context.WithTimeout(ctx, time.Second*5)
	defer cancel()
	result := db.WithContext(dbCtx).Delete(&models.UserInDatabase{}, id)
	return result.Error
}
//class migrations
func CreateClassInDatabase(className string, db *gorm.DB, ctx context.Context) error {
	dbCtx, cancel := context.WithTimeout(ctx, time.Second*3)
	defer cancel()
	classId:=strconv.FormatUint(uint64(uuid.New().ID()),10)
	// Create a table specific to the class
	//tableName := "class_" + className
	sql := `INSERT INTO class_in_database(
        id ,name) VALUES (`+classId+`,"`+className+`");`

	result := db.WithContext(dbCtx).Exec(sql)
	return result.Error
}
func ReadClass(className string, db *gorm.DB, ctx context.Context) ([]models.UserInDatabase, error) {
	dbCtx, cancel := context.WithTimeout(ctx, time.Second*10)
	defer cancel()
	var class models.ClassInDatabase
	var students []models.UserInDatabase
	classId := db.WithContext(dbCtx).Table("class_in_database").Where("name=?",className).Find(&class)
	if classId.Error != nil {
		return nil, classId.Error
	}
	result := db.WithContext(dbCtx).Where("class = ?", class.ClassID).Find(&students)
	if result.Error != nil {
		return nil, result.Error
	}
	return students, nil
}
func DeleteClassFromDatabase(className string, db *gorm.DB, ctx context.Context) error {
	dbCtx, cancel := context.WithTimeout(ctx, time.Second*5)
	defer cancel()
	//tableName := "class_" + className
	sql := `DELETE FROM class_in_database WHERE name = "`+className+`";`
	result := db.WithContext(dbCtx).Exec(sql)
	return result.Error
}
func UpdateClassInDatabase(oldClassName string, newClassName string, db *gorm.DB, ctx context.Context) error {
	dbCtx, cancel := context.WithTimeout(ctx, time.Second*5)
	defer cancel()
	sql:= `UPDATE class_in_database SET name = "`+newClassName+`" WHERE name = "`+oldClassName+`";`
	result := db.WithContext(dbCtx).Exec(sql)
	return result.Error
}
//student migrations
func CreateStudentInDatabase(student models.StudentInDatabase, db *gorm.DB, ctx context.Context) error {
	dbCtx, cancel := context.WithTimeout(ctx, time.Second*3)
	defer cancel()
	result := db.WithContext(dbCtx).Create(&student)
	return result.Error
}
func ReadStudent(classID uint, db *gorm.DB, ctx context.Context) error{
	dbCtx, cancel := context.WithTimeout(ctx, time.Second*10)
	defer cancel()
	var students []models.StudentInDatabase
	result := db.WithContext(dbCtx).Where("class = ?", classID).Find(&students)
	return result.Error
}
func DeleteStudentFromDatabase(studentID uint, db *gorm.DB, ctx context.Context) error {
	dbCtx, cancel := context.WithTimeout(ctx, time.Second*5)
	defer cancel()
	result := db.WithContext(dbCtx).Delete(&models.StudentInDatabase{}, studentID)
	return result.Error
}
func UpdateStudentInDatabase(student models.StudentInDatabase, db *gorm.DB, ctx context.Context) error {
	dbCtx, cancel := context.WithTimeout(ctx, time.Second*5)
	defer cancel()
	result := db.WithContext(dbCtx).Save(&student)
	return result.Error
}