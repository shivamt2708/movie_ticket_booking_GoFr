package Components14

import (
	"gofr.dev/pkg/gofr"
	models "server/models"
)

func Movie(ctx *gofr.Context) (interface{}, error) {
	var customers []models.Movie

	// Getting the customer data from the database using SQL
	rows, err := ctx.DB().QueryContext(ctx.Request().Context(), "SELECT * FROM movies")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var customer models.Movie
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
}