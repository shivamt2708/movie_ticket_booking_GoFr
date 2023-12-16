package main

import (
	"gofr.dev/pkg/gofr"
	Components "server/Components"
	Components1 "server/Components/Components1"
	Components2 "server/Components/Components2"
	Components3 "server/Components/Components3"
	Components4 "server/Components/Components4"
	Components5 "server/Components/Components5"
	Components6 "server/Components/Components6"
	Components7 "server/Components/Components7"
	Components8 "server/Components/Components8"
	Components9 "server/Components/Components9"
	Components10 "server/Components/Components10"
	Components11 "server/Components/Components11"
	Components12 "server/Components/Components12"
	Components13 "server/Components/Components13"
	Components14 "server/Components/Components14"
	Components15 "server/Components/Components15"
	Components16 "server/Components/Components16"
)

func main() {
	app := gofr.New()

	app.POST("/signup/{email}/{username}/{password}/{role}/{location}", Components.Signup)

	app.POST("/add-movie-hall/{email}/{total_seats}/{price}/{name}", Components1.AddMovieHall)

	app.POST("/add-show/{email}/{movie_name}/{hall_name}/{date}/{time}", Components2.AddShow)

	app.POST("/add-movie/{movie_name}", Components3.AddMovie)

	app.POST("/admin/book-ticket/{show_id}/{user_email}", Components4.AdminBookTicket)

	app.PUT("/admin/book-ticket2/{show_id}/{seats_left}", Components5.AdminBookTicket2)

	app.GET("/customer", Components6.Customer)

	app.GET("/my-halls/{email}", Components7.MyHalls)

	app.GET("/my-bookings/{email}", Components8.MyBookings)

	app.GET("/my-shows/{email}/{movie_name}", Components9.MyShows)

	app.GET("/my-shows5/{id}", Components13.MyShows5)

	app.GET("/my-shows3/{email}/{movie_name}/{date}", Components11.MyShows3)

	app.GET("/my-shows4/{email}/{movie_name}/{date}/{time}", Components12.MyShows4)

	app.GET("/my-shows2/{email}/{movie_name}/{hall_name}/{date}/{time}", Components10.MyShows2)

	app.GET("/movie", Components14.Movie)

	app.GET("/{email}/movie", Components15.Movie2)

	app.POST("/login/{email}/{password}", Components16.Login)



	app.Start()
}