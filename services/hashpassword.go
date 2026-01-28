package services
import (
	"golang.org/x/crypto/bcrypt"
)
func HashPassword(password string)(string,error){
	hashedPasswordBytes,err:=bcrypt.GenerateFromPassword([]byte(password),bcrypt.DefaultCost)
	if err!=nil{
		return "",err
	}
	return string(hashedPasswordBytes),nil
}
func CheckPasswordHash(password,hashedPassword string)error{
	err:=bcrypt.CompareHashAndPassword([]byte(hashedPassword),[]byte(password))
	return err
}