package controllers
import (
	"net/http"
	"github.com/AbanoubGirges/malaykaproject/models"
	"encoding/json"
	"github.com/google/uuid"
	"github.com/AbanoubGirges/malaykaproject/services"
)
var userPtr *models.User
func SignupHandler(w http.ResponseWriter,r *http.Request){
	var user models.User
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		services.RespondWithJson(w, 400, struct{}{})
		return
	}
	defer services.RespondWithJson(w, 201, struct{}{})
	user.ID=uuid.New().ID()
	user.Password,_=services.HashPassword(user.Password)
	userPtr=&user
}