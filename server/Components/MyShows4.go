package Components

import (
	"gofr.dev/pkg/gofr"
	models "server/models"
)

func MyShows4(ctx *gofr.Context) (interface{}, error) {
	email := ctx.PathParam("email")
	movie_name := ctx.PathParam("movie_name")
	date := ctx.PathParam("date")
	time := ctx.PathParam("time")

	var customers []models.Show

	// Getting the customer data from the database using SQL
	rows, err := ctx.DB().QueryContext(ctx.Request().Context(), "SELECT * FROM shows where email = ? AND movie_name = ? AND date = ? AND time = ?",email, movie_name, date, time)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var customer models.Show
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
}