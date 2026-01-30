package services
import (
	"github.com/AbanoubGirges/malaykaproject/models"
)
func ToUserInDatabase(user models.User)models.UserInDatabase{
	return models.UserInDatabase{
		ID: user.ID,
		Name: user.Username,
		PhoneNumber: user.PhoneNumber,
		Password: user.Password,
	}
}
func ToStudentInDatabase(student models.Student)models.StudentInDatabase{
	return models.StudentInDatabase{
		ID: student.ID,
		Name: student.Name,
		PhoneNumber: student.PhoneNumber[0],
		Location: student.Location,
		Coordinates: student.Coordinates,
		Age: student.Age,
		Class: student.Class,
		Birthdate: student.Birthdate,
	}
}