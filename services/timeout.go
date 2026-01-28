package services
import (
	"net/http"
)
func RequestTimeout(w http.ResponseWriter, r *http.Request) {
	RespondWithJson(w,504,struct{}{})
}