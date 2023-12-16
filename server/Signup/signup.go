package Signup

import "gofr.dev/pkg/gofr"

func Signup(ctx *gofr.Context) (interface{}, error) {
	email := ctx.PathParam("email")
	username := ctx.PathParam("username")
	password := ctx.PathParam("password")
	role := ctx.PathParam("role")
	location := ctx.PathParam("location")

	// Inserting a customer row in the database using SQL
	data, err := ctx.DB().ExecContext(ctx.Request().Context(),
		"INSERT INTO users (email, username, password, role, location) VALUES (?, ?, ?, ?, ?)",
		email, username, password, role, location)


	return data, err
}