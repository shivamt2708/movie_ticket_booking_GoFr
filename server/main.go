package main

import (
	"gofr.dev/pkg/gofr"
	"fmt"
	Components "server/Components"
)


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

type Show struct {
	Id int `json:"id"`
	Email    string `json:"email"`
	Movie_name string `json:"movie_name"`
	Hall_name     string `json:"hall_name"`
	Seats_left int `json:"seats_left"`
	Date string `json:"date"`
	Time string `json:"time"`
}

type Booking struct {
	Id int `json:"id"`
	ShowID int `json:"show_id"`
	User_email string `json:"user_email"`
}

type Movie struct {
	Id int `json:"id"`
	Movie_name string `json:"movie_name"`
}

func main() {
	app := gofr.New()

	app.POST("/signup/{email}/{username}/{password}/{role}/{location}", Components.Signup)

	app.POST("/add-movie-hall/{email}/{total_seats}/{price}/{name}", Components.AddMovieHall)

	app.POST("/add-show/{email}/{movie_name}/{hall_name}/{date}/{time}", Components.AddShow)

	app.POST("/add-movie/{movie_name}", Components.AddMovie)

	app.POST("/admin/book-ticket/{show_id}/{user_email}", Components.AdminBookTicket)

	app.PUT("/admin/book-ticket2/{show_id}/{seats_left}", Components.AdminBookTicket2)
	
	app.GET("/admin/book-ticket3/{show_id}", func(ctx *gofr.Context) (interface{}, error) {
		show_id := ctx.PathParam("show_id")

		var customers []Show

		// Getting the customer data from the database using SQL
		rows, err := ctx.DB().QueryContext(ctx.Request().Context(), "SELECT * FROM shows where id = ?", show_id)
		if err != nil {
			return nil, err
		}
		defer rows.Close()

		for rows.Next() {
			var customer Show
			if err := rows.Scan(&customer.Id, &customer.Email, &customer.Movie_name, &customer.Hall_name, &customer.Seats_left, &customer.Date, &customer.Time); err != nil {
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


	app.GET("/customer", func(ctx *gofr.Context) (interface{}, error) {
		var customers []Customer

		// Getting the customer data from the database using SQL
		rows, err := ctx.DB().QueryContext(ctx.Request().Context(), "SELECT * FROM users where role = 'user'")
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

	app.GET("/my-bookings/{email}", func(ctx *gofr.Context) (interface{}, error) {
		email := ctx.PathParam("email")
		var customers []Booking

		// Getting the customer data from the database using SQL
		rows, err := ctx.DB().QueryContext(ctx.Request().Context(), "SELECT * FROM bookings WHERE show_id IN (SELECT id FROM shows WHERE email = ?) ORDER BY id",email)
		if err != nil {
			return nil, err
		}
		defer rows.Close()

		for rows.Next() {
			var customer Booking
			if err := rows.Scan(&customer.Id, &customer.ShowID, &customer.User_email); err != nil {
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

	app.GET("/my-shows/{email}/{movie_name}", func(ctx *gofr.Context) (interface{}, error) {
		email := ctx.PathParam("email")
		movie_name := ctx.PathParam("movie_name")
		var customers []Show

		// Getting the customer data from the database using SQL
		rows, err := ctx.DB().QueryContext(ctx.Request().Context(), "SELECT * FROM shows where email = ? AND movie_name = ?",email, movie_name)
		if err != nil {
			return nil, err
		}
		defer rows.Close()

		for rows.Next() {
			var customer Show
			if err := rows.Scan(&customer.Id, &customer.Email, &customer.Movie_name, &customer.Hall_name, &customer.Seats_left, &customer.Date, &customer.Time); err != nil {
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

	app.GET("/my-shows5/{id}", func(ctx *gofr.Context) (interface{}, error) {
		id := ctx.PathParam("id")
		var customers []Show

		// Getting the customer data from the database using SQL
		rows, err := ctx.DB().QueryContext(ctx.Request().Context(), "SELECT * FROM shows where id = ?",id)
		if err != nil {
			return nil, err
		}
		defer rows.Close()

		for rows.Next() {
			var customer Show
			if err := rows.Scan(&customer.Id, &customer.Email, &customer.Movie_name, &customer.Hall_name, &customer.Seats_left, &customer.Date, &customer.Time); err != nil {
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

	app.GET("/my-shows3/{email}/{movie_name}/{date}", func(ctx *gofr.Context) (interface{}, error) {
		email := ctx.PathParam("email")
		movie_name := ctx.PathParam("movie_name")
		date := ctx.PathParam("date")

		var customers []Show

		// Getting the customer data from the database using SQL
		rows, err := ctx.DB().QueryContext(ctx.Request().Context(), "SELECT * FROM shows where email = ? AND movie_name = ? AND date = ?",email, movie_name, date)
		if err != nil {
			return nil, err
		}
		defer rows.Close()

		for rows.Next() {
			var customer Show
			if err := rows.Scan(&customer.Id, &customer.Email, &customer.Movie_name, &customer.Hall_name, &customer.Seats_left, &customer.Date, &customer.Time); err != nil {
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

	app.GET("/my-shows4/{email}/{movie_name}/{date}/{time}", func(ctx *gofr.Context) (interface{}, error) {
		email := ctx.PathParam("email")
		movie_name := ctx.PathParam("movie_name")
		date := ctx.PathParam("date")
		time := ctx.PathParam("time")

		var customers []Show

		// Getting the customer data from the database using SQL
		rows, err := ctx.DB().QueryContext(ctx.Request().Context(), "SELECT * FROM shows where email = ? AND movie_name = ? AND date = ? AND time = ?",email, movie_name, date, time)
		if err != nil {
			return nil, err
		}
		defer rows.Close()

		for rows.Next() {
			var customer Show
			if err := rows.Scan(&customer.Id, &customer.Email, &customer.Movie_name, &customer.Hall_name, &customer.Seats_left, &customer.Date, &customer.Time); err != nil {
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

	app.GET("/my-shows2/{email}/{movie_name}/{hall_name}/{date}/{time}", func(ctx *gofr.Context) (interface{}, error) {
		email := ctx.PathParam("email")
		movie_name := ctx.PathParam("movie_name")
		hall_name := ctx.PathParam("hall_name")
		date := ctx.PathParam("date")
		time := ctx.PathParam("time")
		var customers []Show

		// Getting the customer data from the database using SQL
		rows, err := ctx.DB().QueryContext(ctx.Request().Context(), "SELECT * FROM shows where email = ? AND movie_name = ? AND hall_name = ? AND date = ? AND time = ?",email, movie_name, hall_name, date, time)
		if err != nil {
			return nil, err
		}
		defer rows.Close()

		for rows.Next() {
			var customer Show
			if err := rows.Scan(&customer.Id, &customer.Email, &customer.Movie_name, &customer.Hall_name, &customer.Seats_left, &customer.Date, &customer.Time); err != nil {
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

	app.GET("/movie", func(ctx *gofr.Context) (interface{}, error) {
		var customers []Movie

		// Getting the customer data from the database using SQL
		rows, err := ctx.DB().QueryContext(ctx.Request().Context(), "SELECT * FROM movies")
		if err != nil {
			return nil, err
		}
		defer rows.Close()

		for rows.Next() {
			var customer Movie
			if err := rows.Scan(&customer.Id, &customer.Movie_name); err != nil {
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

	app.GET("/{email}/movie", func(ctx *gofr.Context) (interface{}, error) {
		email := ctx.PathParam("email")
		var customers []Show

		// Getting the customer data from the database using SQL
		rows, err := ctx.DB().QueryContext(ctx.Request().Context(), "SELECT DISTINCT movie_name FROM shows where email = ?", email)
		if err != nil {
			return nil, err
		}
		defer rows.Close()

		for rows.Next() {
			var customer Show
			if err := rows.Scan(&customer.Movie_name); err != nil {
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
			} else if data.Role == "super-admin" {
				fmt.Println("super-admin")
				return "super-admin", nil
			}
		}
	
		return nil, nil
	})



	app.Start()
}