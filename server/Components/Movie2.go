package Components

import (
	"gofr.dev/pkg/gofr"
	models "server/models"
)

func Movie2(ctx *gofr.Context) (interface{}, error) {
	email := ctx.PathParam("email")
	var customers []models.Show

	// Getting the customer data from the database using SQL
	rows, err := ctx.DB().QueryContext(ctx.Request().Context(), "SELECT DISTINCT movie_name FROM shows where email = ?", email)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var customer models.Show
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
}