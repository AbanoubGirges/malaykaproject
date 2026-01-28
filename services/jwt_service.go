package services

import (
	"github.com/golang-jwt/jwt/v5"
	"github.com/AbanoubGirges/malaykaproject/models"
)
var SecretKey string
func TakeSecretKey(key string){
	SecretKey=key
}

func GenerateJWT(user models.UserInDatabase, secretKey string) (string, error) {
	claims := jwt.MapClaims{}
	claims["user_id"] = user.ID
	claims["name"] = user.Name
	//claims["phone_number"] = user.PhoneNumber
	claims["role"] = user.Role
	claims["class"] = user.Class
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signedToken, err := token.SignedString([]byte(secretKey))
	if err != nil {
		return "", err
	}
	return signedToken, nil
}
func ValidateJWT(signedToken string, secretKey string) (jwt.MapClaims, error) {
	token, err := jwt.Parse(signedToken, func(token *jwt.Token) (interface{}, error) {
		return []byte(secretKey), nil
	})
	if err != nil {
		return nil, err
	}
	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		return claims, nil
	} else {
		return nil, err
	}
}
