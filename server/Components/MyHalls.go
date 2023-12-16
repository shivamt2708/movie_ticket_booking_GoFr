package Components

import (
	"gofr.dev/pkg/gofr"
	models "server/models"
)

func MyHalls(ctx *gofr.Context) (interface{}, error) {
	email := ctx.PathParam("email")
	var customers []models.Hall

	// Getting the customer data from the database using SQL
	rows, err := ctx.DB().QueryContext(ctx.Request().Context(), "SELECT * FROM halls where email = ?",email)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var customer models.Hall
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
}