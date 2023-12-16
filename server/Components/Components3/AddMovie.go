package Components3

import "gofr.dev/pkg/gofr"

func AddMovie(ctx *gofr.Context) (interface{}, error) {
	movie_name := ctx.PathParam("movie_name")

	// Inserting a customer row in the database using SQL
	data, err := ctx.DB().ExecContext(ctx.Request().Context(),
		"INSERT INTO movies (movie_name) VALUES (?)",
		movie_name)

	return data, err
}