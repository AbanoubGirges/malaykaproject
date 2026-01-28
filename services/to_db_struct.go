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