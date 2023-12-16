package Components5

import "gofr.dev/pkg/gofr"

func AdminBookTicket2(ctx *gofr.Context) (interface{}, error) {
	show_id := ctx.PathParam("show_id")
	seats_left := ctx.PathParam("seats_left")

	// Inserting a customer row in the database using SQL
		data, err := ctx.DB().ExecContext(ctx.Request().Context(),
			"UPDATE shows SET seats_left = ? where id = ?",
			seats_left, show_id)

	return data, err
}