package Components

import "gofr.dev/pkg/gofr"

func AddShow(ctx *gofr.Context) (interface{}, error) {
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
}