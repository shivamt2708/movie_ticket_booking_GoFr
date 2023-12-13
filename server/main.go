package main

import (
	"context"
	"gofr.dev/pkg/gofr"
	"go.mongodb.org/mongo-driver/bson"
)

type User struct {
	email string `json:"email"`
    username string `json:"username"`
    password string `json:"password"`
    role string `json:"role"`
    location string `json:"location"`
}
type sign_response struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
}

func SignupHandler(ctx *gofr.Context) (interface{}, error) {
    email := ctx.PathParam("email")
    username := ctx.PathParam("username")
    password := ctx.PathParam("password")
    role := ctx.PathParam("role")
    location := ctx.PathParam("location")
    // Connect to MongoDB

    // Insert user data
    collection := ctx.MongoDB.Collection("users")
    collection.InsertOne(context.Background(), bson.M{
        "email": email,
        "username": username,
        "password": password,
        "role": role,
        "location": location,
    })

    // Respond with success message
    return true, nil
}

func main() {
    app := gofr.New()
    // Register your MongoDB-connected handlers
    
    app.POST("/signup/{email}/{username}/{password}/{role}/{location}", SignupHandler)

    // Start the server
    app.Start()
}