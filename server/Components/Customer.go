package Components

import (
	"gofr.dev/pkg/gofr"
	models "server/models"
)

func Customer(ctx *gofr.Context) (interface{}, error) {
	var customers []models.Customer

	// Getting the customer data from the database using SQL
	rows, err := ctx.DB().QueryContext(ctx.Request().Context(), "SELECT * FROM users where role = 'user'")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var customer models.Customer
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
}