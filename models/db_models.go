package models
import "gorm.io/gorm"
type UserInDatabase struct {
	gorm.Model
	ID   uint32   
	Name string		`gorm:"uniqueIndex"`
	PhoneNumber string 
	Password string 
	Role     string	
	Class    uint	`gorm:"foreignkey:Class;references:ClassID"`
}
type ClassInDatabase struct {
	gorm.Model
	ClassID   uint32 
	Name string
	//KhademID string
	//StudentID string
}
type StudentInDatabase struct {
	gorm.Model
	ID       uint32   
	Name string 
	PhoneNumber    string 
	Location string 
	Coordinates string 
	Age 	uint   
	Class    uint	`gorm:"foreignkey:Class;references:ClassID"`
	Birthdate string 
}