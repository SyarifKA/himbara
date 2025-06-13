package dtos

type User struct {
	PhoneNumber string `form:"phoneNumber"`
	Bank        string `form:"bank"`
}
