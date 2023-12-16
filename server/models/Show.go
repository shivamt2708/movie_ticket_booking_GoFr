package models

type Show struct {
	Id int `json:"id"`
	Email    string `json:"email"`
	Movie_name string `json:"movie_name"`
	Hall_name     string `json:"hall_name"`
	Seats_left int `json:"seats_left"`
	Date string `json:"date"`
	Time string `json:"time"`
}