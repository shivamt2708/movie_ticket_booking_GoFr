package models

type Hall struct {
	Email    string `json:"email"`
	Total_seats int `json:"total_seats"`
	Price int `json:"price"`
	Name     string `json:"name"`
}