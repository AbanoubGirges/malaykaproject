package models

type User struct {
	ID       uint32   
	Username string `json:"username"`
	PhoneNumber    string `json:"phone_number"`
	Password string `json:"password"`
	Role     uint 
	Class    uint
}
type UserLoginRequest struct {
	PhoneNumber string `json:"phone_number"`
	Password    string `json:"password"`
}