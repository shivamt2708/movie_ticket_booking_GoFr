package main

import (
	"gofr.dev/pkg/gofr"
	Components "server/Components"
)

func main() {
	app := gofr.New()

	app.POST("/signup/{email}/{username}/{password}/{role}/{location}", Components.Signup)

	app.POST("/add-movie-hall/{email}/{total_seats}/{price}/{name}", Components.AddMovieHall)

	app.POST("/add-show/{email}/{movie_name}/{hall_name}/{date}/{time}", Components.AddShow)

	app.POST("/add-movie/{movie_name}", Components.AddMovie)

	app.POST("/admin/book-ticket/{show_id}/{user_email}", Components.AdminBookTicket)

	app.PUT("/admin/book-ticket2/{show_id}/{seats_left}", Components.AdminBookTicket2)

	app.GET("/customer", Components.Customer)

	app.GET("/my-halls/{email}", Components.MyHalls)

	app.GET("/my-bookings/{email}", Components.MyBookings)

	app.GET("/my-shows/{email}/{movie_name}", Components.MyShows)

	app.GET("/my-shows5/{id}", Components.MyShows5)

	app.GET("/my-shows3/{email}/{movie_name}/{date}", Components.MyShows3)

	app.GET("/my-shows4/{email}/{movie_name}/{date}/{time}", Components.MyShows4)

	app.GET("/my-shows2/{email}/{movie_name}/{hall_name}/{date}/{time}", Components.MyShows2)

	app.GET("/movie", Components.Movie)

	app.GET("/{email}/movie", Components.Movie2)

	app.POST("/login/{email}/{password}", Components.Login)



	app.Start()
}