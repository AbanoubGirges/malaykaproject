package models

type User struct {
	ID       uint32   
	Username string `json:"username"`
	PhoneNumber    string `json:"phone_number"`
	Password string `json:"password"`
	Role     string 
	Class    uint
}
type UserLoginRequest struct {
	PhoneNumber string `json:"phone_number"`
	Password    string `json:"password"`
}
type Student struct {
	ID       uint32   
	Name string `json:"name"`
	PhoneNumber    []string `json:"phone_number"`
	Location string `json:"location"`
	Coordinates string `json:"coordinates"`
	Age 	uint   `json:"age"`
	Class    uint	`json:"class"`
	Birthdate string `json:"birthdate"`
}