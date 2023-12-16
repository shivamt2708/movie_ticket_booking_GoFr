package Components1

import "gofr.dev/pkg/gofr"

func AddMovieHall(ctx *gofr.Context) (interface{}, error) {
	email := ctx.PathParam("email")
	total_seats := ctx.PathParam("total_seats")
	price := ctx.PathParam("price")
	name := ctx.PathParam("name")

	// Inserting a customer row in the database using SQL
	data, err := ctx.DB().ExecContext(ctx.Request().Context(),
		"INSERT INTO halls (email, total_seats, price, name) VALUES (?, ?, ?, ?)",
		email, total_seats, price, name)

	return data, err
}