package models
import "gorm.io/gorm"
type AttendaceInDatabase struct {
	gorm.Model
	Blame uint32	//the one who marked him present
	ClassID uint32 `gorm:"foreignkey:ClassID;references:ClassID"`
	Date string		//`gorm:""`
	Student string `json:"student_id"`
	Present bool	`json:"present"`
}