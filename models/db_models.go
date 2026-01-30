package models
import "gorm.io/gorm"
type UserInDatabase struct {
	gorm.Model
	ID   uint32   
	Name string 
	PhoneNumber string 
	Password string 
	Role     uint	
	Class    uint	`gorm:"foreignkey:Class;references:ClassID"`
}
type ClassInDatabase struct {
	gorm.Model
	ClassID   uint32 
	Name string
	//KhademID string
	//StudentID string
}