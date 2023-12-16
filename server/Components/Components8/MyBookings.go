package Components8

import (
	"gofr.dev/pkg/gofr"
	models "server/models"
)

func MyBookings(ctx *gofr.Context) (interface{}, error) {
		email := ctx.PathParam("email")
		var customers []models.Booking

		// Getting the customer data from the database using SQL
		rows, err := ctx.DB().QueryContext(ctx.Request().Context(), "SELECT * FROM bookings WHERE show_id IN (SELECT id FROM shows WHERE email = ?) ORDER BY id",email)
		if err != nil {
			return nil, err
		}
		defer rows.Close()

		for rows.Next() {
			var customer models.Booking
			if err := rows.Scan(&customer.Id, &customer.ShowID, &customer.User_email); err != nil {
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