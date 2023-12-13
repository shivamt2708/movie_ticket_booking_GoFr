package main

import "gofr.dev/pkg/gofr"
import "fmt"


type Customer struct {
	Email    string `json:"email"`
	Username string `json:"username"`
	Password string `json:"password"`
	Role     string `json:"role"`
	Location string `json:"location"`
}

func main() {
	app := gofr.New()

	app.POST("/signup/{email}/{username}/{password}/{role}/{location}", func(ctx *gofr.Context) (interface{}, error) {
		email := ctx.PathParam("email")
		username := ctx.PathParam("username")
		password := ctx.PathParam("password")
		role := ctx.PathParam("role")
		location := ctx.PathParam("location")

		// Inserting a customer row in the database using SQL
		data, err := ctx.DB().ExecContext(ctx.Request().Context(),
			"INSERT INTO users (email, username, password, role, location) VALUES (?, ?, ?, ?, ?)",
			email, username, password, role, location)


		return data, err
	})

	app.GET("/customer", func(ctx *gofr.Context) (interface{}, error) {
		var customers []Customer

		// Getting the customer data from the database using SQL
		rows, err := ctx.DB().QueryContext(ctx.Request().Context(), "SELECT * FROM users")
		if err != nil {
			return nil, err
		}
		defer rows.Close()

		for rows.Next() {
			var customer Customer
			if err := rows.Scan(&customer.Email, &customer.Username, &customer.Password, &customer.Role, &customer.Location); err != nil {
				return nil, err
			}

			customers = append(customers, customer)
		}

		if err := rows.Err(); err != nil {
			return nil, err
		}

		// Return the customer data
		return nil, nil
	})

	app.POST("/login/{email}/{password}", func(ctx *gofr.Context) (interface{}, error) {
		email := ctx.PathParam("email")
		password := ctx.PathParam("password")
	
		// Getting the customer data from the database using SQL
		rows, err := ctx.DB().QueryContext(ctx.Request().Context(), "SELECT * FROM users WHERE email = ? AND password = ?", email, password)
		if err != nil {
			return nil, err
		}
		defer rows.Close()
	
		if rows.Next() {
			var data Customer
			if err := rows.Scan(&data.Email, &data.Username, &data.Password, &data.Role, &data.Location); err != nil {
				return nil, err
			}
	
			if data.Role == "user" {
				fmt.Println("user")
				return "user", nil
			} else if data.Role == "admin" {
				fmt.Println("admin")
				return "admin", nil
			}
		}
	
		return nil, nil
	})
	

	app.Start()
}