package models

type Booking struct {
	Id int `json:"id"`
	ShowID int `json:"show_id"`
	User_email string `json:"user_email"`
}