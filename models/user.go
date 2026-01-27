package models

type User struct {
	ID       uint32   
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
	Role     uint 
}