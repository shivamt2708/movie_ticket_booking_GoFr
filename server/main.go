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

type Hall struct {
	Email    string `json:"email"`
	Total_seats int `json:"total_seats"`
	Price int `json:"price"`
	Name     string `json:"name"`
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

	app.POST("/add-movie-hall/{email}/{total_seats}/{price}/{name}", func(ctx *gofr.Context) (interface{}, error) {
		email := ctx.PathParam("email")
		total_seats := ctx.PathParam("total_seats")
		price := ctx.PathParam("price")
		name := ctx.PathParam("name")

		// Inserting a customer row in the database using SQL
		data, err := ctx.DB().ExecContext(ctx.Request().Context(),
			"INSERT INTO halls (email, total_seats, price, name) VALUES (?, ?, ?, ?)",
			email, total_seats, price, name)

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
		return customers, nil
	})

	app.GET("/my-halls/{email}", func(ctx *gofr.Context) (interface{}, error) {
		email := ctx.PathParam("email")
		var customers []Hall

		// Getting the customer data from the database using SQL
		rows, err := ctx.DB().QueryContext(ctx.Request().Context(), "SELECT * FROM halls where email = ?",email)
		if err != nil {
			return nil, err
		}
		defer rows.Close()

		for rows.Next() {
			var customer Hall
			if err := rows.Scan(&customer.Email, &customer.Total_seats, &customer.Price, &customer.Name); err != nil {
				return nil, err
			}

			customers = append(customers, customer)
		}

		if err := rows.Err(); err != nil {
			return nil, err
		}

		// Return the customer data
		return customers, nil
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