package Components4

import "gofr.dev/pkg/gofr"

func AdminBookTicket(ctx *gofr.Context) (interface{}, error) {
	show_id := ctx.PathParam("show_id")
	user_email := ctx.PathParam("user_email")

	// Inserting a customer row in the database using SQL
		data, err := ctx.DB().ExecContext(ctx.Request().Context(),
			"INSERT INTO bookings (show_id, user_email) VALUES (?, ?)",
			show_id, user_email)

	return data, err
}