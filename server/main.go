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

	app.POST("/add-show/{email}/{movie_name}/{hall_name}/{date}/{time}", func(ctx *gofr.Context) (interface{}, error) {
		email := ctx.PathParam("email")
		movie_name := ctx.PathParam("movie_name")
		hall_name := ctx.PathParam("hall_name")
		date := ctx.PathParam("date")
		time := ctx.PathParam("time")

		// Inserting a customer row in the database using SQL
		data, err := ctx.DB().ExecContext(ctx.Request().Context(),
			"INSERT INTO shows (email, movie_name, hall_name, seats_left, date, time) VALUES (?, ?, ?, (SELECT total_seats FROM halls WHERE name = ?), ?, ?)",
			email, movie_name, hall_name, hall_name, date, time)

		return data, err
	})

	app.POST("/add-movie/{movie_name}", func(ctx *gofr.Context) (interface{}, error) {
		movie_name := ctx.PathParam("movie_name")

		// Inserting a customer row in the database using SQL
		data, err := ctx.DB().ExecContext(ctx.Request().Context(),
			"INSERT INTO movies (movie_name) VALUES (?)",
			movie_name)

		return data, err
	})

	app.POST("/admin/book-ticket/{show_id}/{user_email}", func(ctx *gofr.Context) (interface{}, error) {
		show_id := ctx.PathParam("show_id")
		user_email := ctx.PathParam("user_email")

		// Inserting a customer row in the database using SQL
		data, err := ctx.DB().ExecContext(ctx.Request().Context(),
			"INSERT INTO bookings (show_id, user_email) VALUES (?, ?)",
			show_id, user_email)

		ctx.DB().ExecContext(ctx.Request().Context(),
			"UPDATE shows SET seats_left = (SELECT total_seats FROM halls WHERE name = (SELECT hall_name FROM shows WHERE id = ?))-1 WHERE id = ?",show_id,show_id)

		return data, err
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
		rows, err := ctx.DB().QueryContext(ctx.Request().Context(), "SELECT * FROM bookings WHERE show_id IN (SELECT id FROM shows WHERE email = ?)",email)
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