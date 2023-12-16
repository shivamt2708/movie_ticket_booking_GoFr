package Components16

import (
	"gofr.dev/pkg/gofr"
	"fmt"
	models "server/models"
)

func Login(ctx *gofr.Context) (interface{}, error) {
	email := ctx.PathParam("email")
	password := ctx.PathParam("password")

	// Getting the customer data from the database using SQL
	rows, err := ctx.DB().QueryContext(ctx.Request().Context(), "SELECT * FROM users WHERE email = ? AND password = ?", email, password)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	if rows.Next() {
		var data models.Customer
		if err := rows.Scan(&data.Email, &data.Username, &data.Password, &data.Role, &data.Location); err != nil {
			return nil, err
		}

		if data.Role == "user" {
			fmt.Println("user")
			return "user", nil
		} else if data.Role == "admin" {
			fmt.Println("admin")
			return "admin", nil
		} else if data.Role == "super-admin" {
			fmt.Println("super-admin")
			return "super-admin", nil
		}
	}

	return nil, nil
}